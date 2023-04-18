# MyGram - Final Project

Api Documentation

## API

The API example app id discribe below

## User API

### Register User
>POST /users/register
```JavaScript
http://localhost:8080/users/register
```

#### Body
```form
username  : username
email     : email@mail.com
password  : password
age       : 20
```

#### Response
```JSON
{
    "data": {
        "id": 6,
        "created_at": "2023-04-18T12:42:44.9362415+07:00",
        "updated_at": "2023-04-18T12:42:44.9362415+07:00",
        "username": "username",
        "email": "email@mail.com",
        "age": 20
    },
    "message": "Data create successfully!"
}
```
##

### Login User
>POST /users/login
```JavaScript
http://localhost:8080/users/login
```

#### Body
```form
email     : email@mail.com
password  : password
```

#### Response
```JSON
{
    "message": "Login successfully!",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYWlsQG1haWwuY29tIiwiaWQiOjZ9.l8Kk7up86fN-mkSfYTDVTzvV-G5QMcSwhIL3V7qSRig"
}
```
##

## Photo API

### Create Photo
>POST /photo
```JavaScript
http://localhost:8080/photo
```

#### Body
```form
title       : Title
photo_url   : https://photo_url.com
caption     : the caption
```
#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": {
        "id": 9,
        "created_at": "2023-04-18T12:52:23.9232127+07:00",
        "updated_at": "2023-04-18T12:52:23.9232127+07:00",
        "title": "Title",
        "caption": "The caption",
        "photo_url": "https://photo_url.com",
        "UserID": 6,
        "User": null,
        "comment": null
    },
    "message": "Photo create successfully!"
}
```
##

### Get One Photo
>GET /photo/:id
```JavaScript
http://localhost:8080/photo/9
```

#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": {
        "id": 9,
        "created_at": "2023-04-18T12:52:23.923212+07:00",
        "updated_at": "2023-04-18T12:52:23.923212+07:00",
        "title": "Title",
        "caption": "The caption",
        "photo_url": "https://photo_url.com",
        "UserID": 6,
        "User": {
            "id": 6,
            "created_at": "2023-04-18T12:42:44.936241+07:00",
            "updated_at": "2023-04-18T12:42:44.936241+07:00",
            "username": "username",
            "email": "email@mail.com",
            "age": 20
        },
        "comment": [
            {
                "id": 6,
                "created_at": "2023-04-18T13:06:21.477616+07:00",
                "updated_at": "2023-04-18T13:06:21.477616+07:00",
                "UserID": 1,
                "photo_id": 9,
                "message": "ngngkrong di sini memang asik",
                "User": null,
                "Photo": null
            },
            {
                "id": 7,
                "created_at": "2023-04-18T13:19:47.183392+07:00",
                "updated_at": "2023-04-18T13:19:47.183392+07:00",
                "UserID": 2,
                "photo_id": 9,
                "message": "Wow Keren",
                "User": null,
                "Photo": null
            }
        ]
    },
    "message": "Get Photo successfully!"
}
```
##

### Get All Photo
>GET /photo/all
```JavaScript
http://localhost:8080/photo/all
```

#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": [
        {
            "id": 1,
            "created_at": "2023-04-15T20:39:21.225581+07:00",
            "updated_at": "2023-04-15T20:39:21.225581+07:00",
            "title": "Bukit Bintang",
            "caption": "Bukit bintang di pinggir jurang",
            "photo_url": "https://images.app.goo.gl/8rAXcXH8usKBrLyz9",
            "UserID": 1,
            "User": null,
            "comment": null
        },
        {
            "id": 8,
            "created_at": "2023-04-16T11:57:31.261325+07:00",
            "updated_at": "2023-04-16T11:57:31.261325+07:00",
            "title": "Test",
            "caption": "test",
            "photo_url": "https://images.app.goo.gl/8rAXcXH8usKBrLyz9",
            "UserID": 2,
            "User": null,
            "comment": null
        },
        {
            "id": 9,
            "created_at": "2023-04-18T12:52:23.923212+07:00",
            "updated_at": "2023-04-18T12:52:23.923212+07:00",
            "title": "Title",
            "caption": "The caption",
            "photo_url": "https://photo_url.com",
            "UserID": 6,
            "User": null,
            "comment": null
        }
    ],
    "message": "Get Photos successfully!"
}
```
##

### Update Photo
>PUT /photo/:id
```JavaScript
http://localhost:8080/photo/9
```

#### Body
```form
title       : Title
photo_url   : https://photo_url.com
caption     : The Captions for photo
```
#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": {
        "id": 9,
        "updated_at": "2023-04-18T13:23:04.4175955+07:00",
        "title": "Title",
        "caption": "The Captions for photo",
        "photo_url": "https://photo_url.com",
        "UserID": 6,
        "User": null,
        "comment": null
    },
    "message": "Update Photo successfully!"
}
```
##


### Delete Photo
>DELETE /photo/:id
```JavaScript
http://localhost:8080/photo/9
```
#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "message": "Delete Succesfully"
}
```
##

## Comment API

### Create Comment
>POST /comment
```JavaScript
http://localhost:8080/comment
```

#### Body
```form
message      : message
photo_id     : 9
```

#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": {
        "id": 7,
        "created_at": "2023-04-18T13:19:47.1833928+07:00",
        "updated_at": "2023-04-18T13:19:47.1833928+07:00",
        "UserID": 2,
        "photo_id": 9,
        "message": "Wow Keren",
        "User": null,
        "Photo": null
    },
    "message": "Comment create successfully!"
}
```
##

### Get One Comment
>GET /comment/:id
```JavaScript
http://localhost:8080/comment/7
```

#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": {
        "id": 7,
        "created_at": "2023-04-18T13:19:47.183392+07:00",
        "updated_at": "2023-04-18T13:19:47.183392+07:00",
        "UserID": 2,
        "photo_id": 9,
        "message": "Wow Keren",
        "User": {
            "id": 2,
            "created_at": "2023-04-15T20:37:06.239731+07:00",
            "updated_at": "2023-04-15T20:37:06.239731+07:00",
            "username": "user",
            "email": "user@gmail.com",
            "age": 20
        },
        "Photo": {
            "id": 9,
            "created_at": "2023-04-18T12:52:23.923212+07:00",
            "updated_at": "2023-04-18T13:23:04.417595+07:00",
            "title": "Title",
            "caption": "The Captions for photo",
            "photo_url": "https://photo_url.com",
            "UserID": 6,
            "User": null,
            "comment": null
        }
    },
    "message": "Get Comment successfully!"
}
```
##

### Get All Comment from Photo
>GET /comment/all/:photo_id
```JavaScript
http://localhost:8080/comment/all/9
```

#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": [
        {
            "id": 6,
            "created_at": "2023-04-18T13:06:21.477616+07:00",
            "updated_at": "2023-04-18T13:06:21.477616+07:00",
            "UserID": 1,
            "photo_id": 9,
            "message": "ngngkrong di sini memang asik",
            "User": {
                "id": 1,
                "created_at": "2023-04-15T20:36:36.083083+07:00",
                "updated_at": "2023-04-15T20:36:36.083083+07:00",
                "username": "public",
                "email": "public@gmail.com",
                "age": 20
            },
            "Photo": {
                "id": 9,
                "created_at": "2023-04-18T12:52:23.923212+07:00",
                "updated_at": "2023-04-18T13:23:04.417595+07:00",
                "title": "Title",
                "caption": "The Captions for photo",
                "photo_url": "https://photo_url.com",
                "UserID": 6,
                "User": null,
                "comment": null
            }
        },
        {
            "id": 7,
            "created_at": "2023-04-18T13:19:47.183392+07:00",
            "updated_at": "2023-04-18T13:19:47.183392+07:00",
            "UserID": 2,
            "photo_id": 9,
            "message": "Wow Keren",
            "User": {
                "id": 2,
                "created_at": "2023-04-15T20:37:06.239731+07:00",
                "updated_at": "2023-04-15T20:37:06.239731+07:00",
                "username": "user",
                "email": "user@gmail.com",
                "age": 20
            },
            "Photo": {
                "id": 9,
                "created_at": "2023-04-18T12:52:23.923212+07:00",
                "updated_at": "2023-04-18T13:23:04.417595+07:00",
                "title": "Title",
                "caption": "The Captions for photo",
                "photo_url": "https://photo_url.com",
                "UserID": 6,
                "User": null,
                "comment": null
            }
        }
    ],
    "message": "Get Comments successfully!"
}
```
##

### Update Comment
>PUT /comment/:id
```JavaScript
http://localhost:8080/comment/6
```

#### Body
```form
message       : message
```
#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": {
        "id": 6,
        "updated_at": "2023-04-18T13:35:17.465016+07:00",
        "UserID": 1,
        "photo_id": 0,
        "message": "Di sini memang asik buat ngongkrong",
        "User": null,
        "Photo": null
    },
    "message": "Update Comment successfully!"
}
```
##


### Delete Comment
>DELETE /comment/:id
```JavaScript
http://localhost:8080/comment/6
```
#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "message": "Delete Succesfully"
}
```
##

## Social Media API

### Create Social Media
>POST /socialmedia
```JavaScript
http://localhost:8080/socialmedia
```

#### Body
```form
name                : user_name
social_media_url    : https://www.social_media_url.com
```

#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": {
        "id": 5,
        "created_at": "2023-04-18T13:41:01.7696252+07:00",
        "updated_at": "2023-04-18T13:41:01.7696252+07:00",
        "name": "user_name",
        "social_media_url": "https://www.social_media_url.com",
        "UserID": 6,
        "User": null
    },
    "message": "Social Media Create successfully!"
}
```
##

### Get One Social Media
>GET /socialmedia/:id
```JavaScript
http://localhost:8080/socialmedia/5
```

#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": {
        "id": 5,
        "created_at": "2023-04-18T13:41:01.769625+07:00",
        "updated_at": "2023-04-18T13:41:01.769625+07:00",
        "name": "user_name",
        "social_media_url": "https://www.social_media_url.com",
        "UserID": 6,
        "User": {
            "id": 6,
            "created_at": "2023-04-18T12:42:44.936241+07:00",
            "updated_at": "2023-04-18T12:42:44.936241+07:00",
            "username": "username",
            "email": "email@mail.com",
            "age": 20
        }
    },
    "message": "Get Social Media successfully!"
}
```
##

### Get All Social Media
>GET /socialmedia/all
```JavaScript
http://localhost:8080/socialmedia/all
```

#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": [
        {
            "id": 4,
            "created_at": "2023-04-15T21:06:04.411958+07:00",
            "updated_at": "2023-04-15T21:06:29.239773+07:00",
            "name": "user is user",
            "social_media_url": "https://www.instagram.com/p/Cq90-f_PIi6/?utm_source=ig_web_copy_link",
            "UserID": 2,
            "User": {
                "id": 2,
                "created_at": "2023-04-15T20:37:06.239731+07:00",
                "updated_at": "2023-04-15T20:37:06.239731+07:00",
                "username": "user",
                "email": "user@gmail.com",
                "age": 20
            }
        },
        {
            "id": 5,
            "created_at": "2023-04-18T13:41:01.769625+07:00",
            "updated_at": "2023-04-18T13:41:01.769625+07:00",
            "name": "user_name",
            "social_media_url": "https://www.social_media_url.com",
            "UserID": 6,
            "User": {
                "id": 6,
                "created_at": "2023-04-18T12:42:44.936241+07:00",
                "updated_at": "2023-04-18T12:42:44.936241+07:00",
                "username": "username",
                "email": "email@mail.com",
                "age": 20
            }
        }
    ],
    "message": "Get Social Media successfully!"
}
```
##

### Update Social Media
>PUT /socialmedia/:id
```JavaScript
http://localhost:8080/socialmedia/5
```

#### Body
```form
name                : user name
social_media_url    : https://www.social_media_url.com
```
#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "data": {
        "id": 5,
        "updated_at": "2023-04-18T13:44:01.953716+07:00",
        "name": "user name",
        "social_media_url": "https://www.social_media_url.com",
        "UserID": 6,
        "User": null
    },
    "message": "Update Social Media successfully!"
}
```
##


### Delete Social Media
>DELETE /socialmedia/:id
```JavaScript
http://localhost:8080/socialmedia/5
```
#### Authorization
> Bearer Token
```form
Token : <token>
```

#### Response
```JSON
{
    "message": "Delete Succesfully"
}
```
##

