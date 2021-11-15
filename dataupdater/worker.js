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
const TOKEN = process.env.TOKEN || "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IlNVUEVSQURNSU4iLCJzY29wZSI6Ii91cGRhdGVfZGF0YSIsImlhdCI6MTUxNjIzOTAyMn0.Ey0Z_w7Ans7DxZWOohVrJ2yGQxoT1Eom0F7sl7CYm2M";

const datafolders = [
  'artifacts',
  'bait',
  'characters',
  'character_exp_material',
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
  'weapon_enhancement_material',
  'weapons',
];

// App
const upload = multer({ dest: path.join(__dirname, 'data') })
const app = express();
app.get('/', (req, res) => {
  // Return helloo world
  res.send('Hello World!');
});

app.post('/updateData', upload.array('data'), async (req, res) => {
  if (req.headers.authorization !== `Bearer ${TOKEN}`) {
    console.log('Invalid token');
    res.send('Hello World!');
    return;
  }

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
