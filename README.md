## Homework №2 for Dmitrii Fadeev
# CHGK-Telegram-Bot
#### Available features:
1. Random questions are taken from https://db.chgk.info/
2. A user requests from bot a new question and a timer is started waiting for an answer
3. Database of users who answered correctly, possibility to ask for rating via telegram bot
## Implementation Concept

![implementation concept](docs/concept.drawio.png "Implementation Concept")

Users(developers) will use telegram bot as a communication channel

Any user is able to register in telegram bot.

After registration a user can request for a random question taken from https://db.chgk.info/xml/random API. In parallel the question will be saved to internal service database (PostgeSQL)

Bot will print a random question and will start a timer waiting for an answer, in case of correct answer this will be added to database to keep a rating of users