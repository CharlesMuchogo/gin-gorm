## Geting started

## .env
Add .env in your project root directory with the following format

```POSTGRES_USER=postgres_user```

```POSTGRES_PASSWORD=your_password```

```POSTGRES_DB=your_database```

```HOST=postgres```

```PORT=5432```

### Start Application
1. ```sudo systemctl stop postgresql.service``` - Stop postgres
2. ```docker-compose up --build``` - Build and start application (for initial setup)
3. ```docker-compose up``` - Start app without building (for continuous development)
4. ```docker-compose up -d``` - Start in detachment mode


## Database Management
1. ```docker-compose exec db psql --username=postgres_user --dbname=your_password```
