import React, { useState, useEffect } from "react";
import { createBook, updateBook } from "../services";

const BookForm = ({ bookToEdit, onFormSubmit }) => {
  const [book, setBook] = useState({ ID: "", name: "", author: "", rilis: "" });

  // Mengisi form jika ada buku yang akan diedit
  useEffect(() => {
    if (bookToEdit) {
      setBook(bookToEdit);
    }
  }, [bookToEdit]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setBook({ ...book, [name]: value });
  };

  const handleAdd = (e) => {
    createBook(book).then(() => onFormSubmit());
  };

  const handleEdit = (e) => {
    e.preventDefault();
    updateBook(book.ID, book).then(() => onFormSubmit());
  };

  return (
    <form>
      <h2>Manage Book</h2>

      {/* Form untuk ID (optional saat tambah atau edit) */}
      <input
        type="number"
        name="ID"
        value={book.ID || ""}
        onChange={handleChange}
        placeholder="Book ID (optional)"
      />
      <input
        type="text"
        name="name"
        value={book.name}
        onChange={handleChange}
        placeholder="Book Name"
      />
      <input
        type="text"
        name="author"
        value={book.author}
        onChange={handleChange}
        placeholder="Author"
      />
      <input
        type="text"
        name="rilis"
        value={book.rilis}
        onChange={handleChange}
        placeholder="Release Date"
      />

      {/* Tombol Add dan Edit */}
      <div style={{ display: "flex", gap: "10px" }}>
        <button type="submit" onClick={handleAdd}>
          Add
        </button>

        <button type="submit" onClick={handleEdit}>
          Edit
        </button>
      </div>
    </form>
  );
};

export default BookForm;
