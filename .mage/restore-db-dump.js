// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/* eslint-disable no-console */
/* eslint-disable import/no-commonjs */

const fs = require('fs')
const { Client } = require('pg')
const yargs = require('yargs')

const argv = yargs.argv

const dbName = argv.db || 'ttn_lorawan_test'

const client = new Client({
  user: 'root',
  host: 'localhost',
  database: dbName,
  port: 26257,
})
client.connect()
const sql = fs.readFileSync('.cache/sqldump.sql', 'utf8')
client.query('DROP DATABASE ttn_lorawan_test', (err, res) => {
  client.query('CREATE DATABASE ttn_lorawan_test', (err, res) => {
    client.query(sql, (err, res) => {
      // console.log(err, res)
      client.end()
    })
  })
})
