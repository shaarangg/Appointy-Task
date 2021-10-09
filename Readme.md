# Appointy-Instagram-clone-API

# About The Project
This project was given as a Task for the 2 month internship program at Appointy. The project is built using go and MongoDB.


## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

Install global dependencies
* go
* MongoDB or a Mongo Cluster

### Installation
 
1. Clone the repo 
```sh
https://github.com/shaarangg/Appointy-Task.git
cd Appointy-Task
```
2. Install dependencies
```sh
go mod download
```

3. Start the project
```sh
go run main.go
```

## API Endpoints

### 1. Add a user to the user collection
```http
  POST /user
```

| Description                     | Type     | Parameter | 
| :------------------------------:| :-------:| :--------: |
|  Name | `string` | `body`    |
|  Email | `string` | `body`    | 
|  Password   |  `string`|  `body`   |


### Response format

```json
{
    "InsertedID": "61620f2b4b5c7357a6674f11"
}
```

### 2. Get a user by the id
```http
  GET /posts/
```

### Response format

```json
{
    "ID": "6161f2e3e6cc8dbf2a559459",
    "Name": "Shaarang Singh",
    "Email": "shaarangg.singh@gmail.com"
}
```
### 3. Add a post to the post collection
```http
  POST /student/:sid
```
| Description                     | Type     | Parameter | 
| :------------------------------:| :-------:| :--------: | 
|  Uid | `ObjectID` | `body`    | 
|  Caption | `string` | `body`    |
|  Image_url   |  `string`|  `body`   |
|  Timestamp   |  `string`|  `body`   |

### Response format

```json
{
    "InsertedID": "61620f2b4b5c7357a6674f11"
}
```

### 4. Get a post by the id
```http
  GET /posts/:id
```

### Response format

```json
{
    "ID": "6161f5dc37246212fa47a41a",
    "Uid": "6161f5041c894b4f8a83f582",
    "Caption": "My first Post",
    "Image_url": "google.com/images",
    "Timestamp": "2021-10-10 01:34:44.49766 +0530 IST m=+18.081275601"
}
```

### 5. Get all the posts of a user
```http
  GET /users/posts/:id
```
### Response format

```json
[
    {
        "ID": "6161f5dc37246212fa47a41a",
        "Uid": "6161f5041c894b4f8a83f582",
        "Caption": "My first Post",
        "Image_url": "google.com/images",
        "Timestamp": "2021-10-10 01:34:44.49766 +0530 IST m=+18.081275601"
    },
    {
        "ID": "6161f5f237246212fa47a41b",
        "Uid": "6161f5041c894b4f8a83f582",
        "Caption": "My second Post",
        "Image_url": "google.com/images",
        "Timestamp": "2021-10-10 01:35:06.6703448 +0530 IST m=+40.253960401"
    },
    {
        "ID": "6161f5f637246212fa47a41c",
        "Uid": "6161f5041c894b4f8a83f582",
        "Caption": "My third Post",
        "Image_url": "google.com/images",
        "Timestamp": "2021-10-10 01:35:10.5058425 +0530 IST m=+44.089458101"
    },
    {
        "ID": "6161f5fb37246212fa47a41d",
        "Uid": "6161f5041c894b4f8a83f582",
        "Caption": "My fourth Post",
        "Image_url": "google.com/images",
        "Timestamp": "2021-10-10 01:35:15.8795415 +0530 IST m=+49.463157101"
    }
]
```
## Testing
```
go test -v
```