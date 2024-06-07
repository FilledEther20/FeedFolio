require('dotenv').config();
const express = require('express');
const mongoose = require('mongoose');
const cors = require('cors');
const PORT = process.env.PORT || 3000;
const app = express();
app.use(express.json());
app.use(cors());
app.listen(PORT, () => {
	console.log(`Listenning on port : ${PORT}`);
});
