## Comment Struct

```json
User{
    "id": string,
    "email": string,
    "password": string,
    "firstName":string,
    "lastName":string,
    "image":string,
}

UserPublicData{
    "id":string,
    "email":string,
    "firstName":string,
    "lastName":string,
    "image":string,
}
```

## GetUsersPublic

- path: `/user/public`
- method: `GET`
- returns:
    ```js
    {
        "users": UserPublicData[],
        "success": Number,
        "message" : String
    }
    ```


## Get

- path : `/user/single/{id}`
- method: `GET`
- returns:
    ```js
    {
        "user": User,
        "success": Number,
        "message" : String
    }
    ```


## Create
- path: `/user/create`
- method: `POST`
- params:
   * Header:
  
        |  Name | Description                           | Type   |
        |:---------:|---------------------------------------|--------|
        | x-auth-token | authentication token of the user  | String |
    ```js
    {
        "email": string,
        "password": string,
        "firstName":string,
        "lastName":string,
    }
    ```

- returns:
    ```js
    {
        "user": NULL,
        "success": Number,
        "message" : String
    }

## Update 
- path: `/user/update`
- method: `PUT`
- params:
   * Header:
  
        |  Name | Description                           | Type   |
        |:---------:|---------------------------------------|--------|
        | x-auth-token | authentication token of the user  | String |
    ```js
    {
        "id": string,
        "email": string,
        "password": string,
        "firstName":string,
        "lastName":string,
        "image":string,
    }
    ```
- returns:
    ```js
    {
        "user": NULL,
        "success": Number,
        "message" : String
    }