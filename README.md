## Bybit-bot
Based on telegram signal channels
``` Go
go mod tidy
go run main.go
```
---
## Only for Educational Purposes
Please Trading is dangerous and more if you are letting a bot doing it for you. 

---
Model for channel 1:
``` Channel 1
🔥 #SSV/USDT (Long📈, x20) 🔥

Entry - 44.08
Take-Profit:

🥉 44.9796 (40% of profit)
🥈 45.4433 (60% of profit)
🥇 45.9167 (80% of profit)
🚀 46.4 (100% of profit)
```
Model for channel 2:
``` Channel 2
Long/Buy #SEI/USDT  

Enter Point - 6670

Targets: 6705 - 6735 - 6770 - 6805

Leverage - 10x

Stop Loss - 6400
```
---
## .env FILE
```
BYBIT_ENDPOINT=https://api.bybit.com
BYBIT_API_KEY=your_bybit_key
BYBIT_SECRET_KEY=yout_bybit_secret

APP_ID=your_telegram_id
APP_HASH=your_telegram_hash
SESSION_DIR=.\src\config\telegram\log

DATABASE_USER=your_postgres_user
DATABASE_PWD=your_postgres_password
DATABASE_PORT=your_postgres_url_port
DATABASE_NAME=your_database_name
DATABASE_URL=your_database_url_connection

```
---
## Config
1. `Bybit`:
   API to do automatic trading Bybit-go.
3. `Postgres`:
   Database for saving orders and positions. You can found the tables and functions on `SQL` folder.
   Be sure to use dbo_trading_bot as principal schema.
5. `Telegram`:
   Gotd as Telegram client to recieve real time messages and configure the orders on bybit.
7. `Server`:
   Our webserver to have everything active
