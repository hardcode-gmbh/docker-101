const express = require('express')
const app = express()
const port = 3000

const DEFAULT_NAME = process.env.DEFAULT_NAME;

app.get('/', (req, res) => {
  res.send('Hello World!')
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})