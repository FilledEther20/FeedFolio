require('dotenv').config();
const express = require('express');
const mongoose = require('mongoose');
const cors = require('cors');
const PORT = process.env.PORT || 3000;
const app = express();
const feedRoutes=require('./routes/routes')
app.use(express.json());
app.use(cors());

app.use('/',feedRoutes)


app.listen(PORT, () => {
	console.log(`Listenning on port : ${PORT}`);
});
