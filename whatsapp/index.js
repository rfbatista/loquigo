const mongoose = require('mongoose');
const logger = require('pino')()
const venom = require('venom-bot');
const mark = require('./mark')
require('dotenv').config()

venom
	.create({
		session: 'session-name', //name of session
		multidevice: false // for version not multidevice use false.(default: true)
	})
	.then((client) => start(client))
	.catch((erro) => {
		console.log(erro);
	});


const sendMessage = (client) => {
	client
		.sendText('16465937900@c.us', 'Oi')
		.then((result) => {
			console.log('Result: ', result);
			mark.start()
		})
		.catch((erro) => {
			console.error('Error when sending: ', erro); //return object error
		});
}

const receiveMessage = async (client) => {
	await mark.end()
  await new Promise(r => setTimeout(r, Math.floor(Math.random() * 2000) + 1000));
	sendMessage(client)
}

const URI = process.env.MONGO_URI

const start = async (client) => {
	await mongoose.connect(URI);
	client.onMessage((message) => receiveMessage(client));
	sendMessage(client)
}