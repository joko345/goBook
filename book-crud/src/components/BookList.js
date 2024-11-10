import React, { useEffect, useState } from "react";
import { getBooks, deleteBook } from "../services";

const BookList = ({ onEdit }) => {
  const [books, setBooks] = useState([]);

  useEffect(() => {
    getBooks().then((response) => setBooks(response.data));
  }, []);

  const handleDelete = (id) => {
    deleteBook(id).then(() => setBooks(books.filter((book) => book.ID !== id)));
  };

  return (
    <div>
      <h2>Book List</h2>
      <ul>
        {books.map((book) => (
          <li key={book.ID}>
            {book.Name} by {book.Author} - {book.Rilis}
            <button onClick={() => onEdit(book)}>Edit</button>
            <button onClick={() => handleDelete(book.ID)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default BookList;
