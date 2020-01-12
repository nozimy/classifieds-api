# Classifieds JSON API

- собрать контейнер `sudo docker build -t nozimdev/classifieds-api .`
- запустить контейнер `sudo docker run -p 5000:5000 --name classifieds-api -t nozimdev/classifieds-api`
- остановить работу контейнера `docker stop classifieds-api`
- удалить контейнер `docker rm classifieds-api`

### GET /api/items
 
Метод получения списка объявлений

Параметры: 

- `sort` : date | price
- `desc`: true | false
- `page`: [0-9]+

Пример запроса: 

```
curl -X GET \
  'http://localhost:5000/api/items?desc=true&sort=price&page=1' \
  -H 'Accept: */*' \
  -H 'Accept-Encoding: gzip, deflate' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Host: localhost:5000' \
  -H 'cache-control: no-cache'
```

### GET /api/item/:id 

Метод получения конкретного объявления

Параметры:

- `fields`: [description,photos]

Пример запроса: 

```
curl -X GET \
  'http://localhost:5000/api/item/1?fields=description,photos' \
  -H 'Accept: */*' \
  -H 'Accept-Encoding: gzip, deflate' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Host: localhost:5000' \
  -H 'cache-control: no-cache'
```

### POST /items 

Метод создания объявления

Body:

- name: string
- price: float
- photos: string[]
- description: string

Пример запроса: 

```
curl -X POST \
  http://localhost:5000/api/items \
  -H 'Accept: */*' \
  -H 'Accept-Encoding: gzip, deflate' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Length: 185' \
  -H 'Content-Type: application/json' \
  -H 'Host: localhost:5000' \
  -H 'cache-control: no-cache' \
  -d '{
	"name":"продается квартира 1м 23",
	"description":"писание продается квартира 1м",
	"price": 100,
	"photos": ["image1.png", "image2.png"]
}'
```

