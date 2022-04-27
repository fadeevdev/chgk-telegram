## Homework №2 for Dmitrii Fadeev
# Gitlab-Telegram-Bot
#### Available features:
1. Register developer, level and vacation
2. Create teams of developers
3. Join teams or standalone developers to project groups or separate project in Gitlab
4. Assign one or more reviewers (configuration should be available in case of more, because of limits existing in free Gitlab versionfor just one reviewer)
5. Notifications via messenger about new tasks
6. "Assign to me button" for re-assigning a developer
7. Registration of developers via messenger
## Implementation Concept

![implementation concept](docs/concept.png "Implementation Concept")

Users(developers) will use telegram bot as a communication channel
Telegram bot will use HTTP proxy to gRPC service, which will be connected to Gitlab API and PostgreSQL Database