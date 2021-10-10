'use strict';

const express = require('express');
const path = require('path');
const fs = require('fs');
const { MongoClient } = require("mongodb");

// Constants
const PORT = 3000;
const HOST = '0.0.0.0';
const DATA_DIR = path.join(__dirname, 'data');
const DB_USER = encodeURIComponent("root");
const DB_PASS = encodeURIComponent("example");
const DB_HOST = "localhost";

const uri = `mongodb://${DB_HOST}`;

const client = new MongoClient(uri, { auth: { username: DB_USER, password: DB_PASS } });

// App
const app = express();
app.post('/updateData', async (req, res) => {
  try {
    await client.connect();

    // for await
    for await (const file of fs.readdirSync(DATA_DIR)) {
      const filePath = path.join(DATA_DIR, file);
      const [_, language, ...names] = file.split('_');
      const name = names.join('_').replace('.min.json', '').trim();

      const dbo = client.db(`genshindata_${language}`);
      try {
        await dbo.dropCollection(name);
      } catch (e) {
        // ignore
      }
      const collection = await dbo.createCollection(name);
      await collection.insertMany(JSON.parse(fs.readFileSync(filePath)));
      console.log(`[${language}] Collection ${name} created.`);
    }

    res.send('Hello World');

  } finally {
    await client.close();
  }
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
