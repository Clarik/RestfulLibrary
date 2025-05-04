### RestfulLibrary

A RESTful API for managing a library system built with Go, Gin, and PostgreSQL.

## How to Run

Prerequisites
* Install Docker and Docker Compose on your machine.
* Ensure docker-compose.yml and Dockerfile are properly configured.

Steps to Run
1. **Start the services**:
   ```
   docker-compose up -d
   ```
2. **Verify the services**:
    * Check if the application is running:
      ```
      curl http://localhost:8080
      ```
      You should see a response or a 404 error (indicating the server is running).
    * Check if PostgreSQL is running:
      ```
      docker logs postgres
      ```
      Look for database system is ready to accept connections.
3. **Stop the Services**:
   ```
   docker-compose down
   ```

## Run Without docker
  Run the following command to run the library:
  ```
  go run .
  ```

## How to Automated Test
  Automated Tests
  Run the following command to execute the automated tests:
  ```
  go test ./automated-test
  ```

## Manual Testing
  You can use the ```.http``` files in the api-test folder to test the API endpoints using tools like VS Code REST Client or Postman.

## Steps to Manual Test with VS Code REST Client:
  1. Open any ```.http``` file (e.g., ```create-book.http```).
  2. Click the "Send Request" button above the HTTP request.
  3. View the response in the editor.

## Example Manual Requests
You can find example requests in the api-test folder:

* ```create-book.http```: Create a book.
* ```get-book.http```: Get all books.
* ```get-book-pagination.http```: Get books with pagination.
* ```update-book.http```: Update a book.
* ```delete-book.http```: Delete a book.


## API Documentation
  Base URL
  ```
  http://localhost:8080
  ```
  Endpoints
  1. **Create a Book**
  * **URL**: ```/books```
  * **Method**: ```POST```
  * **Headers**:
    * ```Content-Type: application/json```
    * ```Authorization: static-token```
  * **Body**:
    ```
    {
    "title": "Book Title",
    "author": "Author Name",
    "description": "Book Description"
    }
    ```
  * **Response**:
    * **201** Created:
      ```
      {
        "message": "Book created",
        "book": {
          "id": 1,
          "title": "Book Title",
          "author": "Author Name",
          "description": "Book Description"
        }
      }
      ```
    * **400 Bad Request**: Missing or invalid fields.
    * **401 Unauthorized**: Missing or invalid token.
  ---
  2. **Get All Books**
  * **URL**: ```/books```
  * **Method**: ```GET```
  * **Headers**:
    * ```Content-Type: application/json```
  * **Response**:
    * **200 OK**:
    ```
    {
      "message": "Get Books Data",
      "books": [
        {
          "id": 1,
          "title": "Book Title",
          "author": "Author Name",
          "description": "Book Description"
        }
      ]
    }
    ```
    * **500 Internal Server Error**: Database issues.
  ---
  3. **Get Books with Pagination**
  * **URL**: ```/books?page={page}&size={size}```
  * **Method**: ```GET```
  * **Headers**:
    * ```Content-Type: application/json```
  * **Query Parameters**:
    * ```page```: Page number (default: 1).
    * ```size```: Number of items per page (default: 5).
  * **Response**:
    * **200 OK**: Paginated list of books.
    * **400 Bad** Request: Invalid query parameters.
  ---
  4. **Update a Book**
  * **URL**: ```/books/{id}```
  * **Method**: ```PUT```
  * **Headers**:
    * ```Content-Type: application/json```
    * ```Authorization: static-token```
  * **Body**:
    ```
    {
      "title": "Updated Title",
      "author": "Updated Author",
      "description": "Updated Description"
    }
    ```
  * **Response**:
    * **200 OK**:
      ```
      {
        "message": "Book Updated Successfully",
        "book": {
          "id": 1,
          "title": "Updated Title",
          "author": "Updated Author",
          "description": "Updated Description"
        }
      }
      ```
    * **400 Bad Request**: Invalid fields or body.
    * **401 Unauthorized**: Missing or invalid token.
    * **404 Not Found**: Book not found.
  ---
  5. **Delete a Book**
  * **URL**: ```/books/{id}```
  * **Method**: ```DELETE```
  * **Headers**:
    * ```Authorization: static-token```
  * **Response**:
    * **200 OK**:
      ```
      {
        "message": "Book removed successfully"
      }
      ```
    * **401 Unauthorized**: Missing or invalid token.
    * **404 Not Found**: Book not found.


