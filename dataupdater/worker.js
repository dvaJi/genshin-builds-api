'use strict';

const express = require('express');
const path = require('path');
const fs = require('fs');
const multer = require('multer');
const { MongoClient } = require("mongodb");

// Constants
const PORT = 3000;
const HOST = '0.0.0.0';
const DB_USER = encodeURIComponent(process.env.DATABASE_USER || "root");
const DB_PASS = encodeURIComponent(process.env.DATABASE_PASSWORD || "example");
const DB_HOST = process.env.DATABASE_HOST || "localhost";

const uri = `mongodb://${DB_HOST}`;

const client = new MongoClient(uri, { auth: { username: DB_USER, password: DB_PASS } });

const datafolders = [
  'artifacts',
  'bait',
  'characters',
  'common_materials',
  'elemental_stone_materials',
  'fish',
  'fishing_rod',
  'food',
  'ingredients',
  'jewels_materials',
  'local_materials',
  'potions',
  'talent_lvl_up_materials',
  'weapon_primary_materials',
  'weapon_secondary_materials',
  'weapons',
];

// App
const upload = multer({ dest: path.join(__dirname, 'data') })
const app = express();
app.post('/updateData', upload.array('data', 13), async (req, res) => {
  try {
    await client.connect();

    for await (const file of req.files) {
      const json = JSON.parse(fs.readFileSync(file.path).toString());
      const language = file.originalname.split('.')[0].replace('data_', '');

      const dbo = client.db(`genshindata_${language}`);

      for await (const folder of datafolders) {
        const data = json[folder];

        try {
          await dbo.dropCollection(folder);
        } catch (e) {
          // ignore
        }

        const collection = await dbo.createCollection(folder);
        await collection.insertMany(data);
        console.log(`[${language}] Collection ${folder} created.`);
      }

      // Delete file after insert to db
      fs.unlinkSync(file.path);
    }

  } catch (err) {
    res.send(err);
  } finally {
    await client.close();
  }

  res.send('Hello World');
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
