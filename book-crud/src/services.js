import axios from "axios";

const API_URL = "http://localhost:8080/book/";

export const getBooks = () => axios.get(API_URL);
export const getBookById = (id) => axios.get(`${API_URL}${id}`);
export const createBook = (book) => axios.post(API_URL, book);
export const updateBook = (id, book) =>
  axios.put(`${API_URL}${id}`, {
    name: book.name,
    author: book.author,
    rilis: book.rilis,
  });
export const deleteBook = (id) => axios.delete(`${API_URL}${id}`);
