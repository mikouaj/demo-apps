package dev.stefaniak.demo.service;

import dev.stefaniak.demo.db.BookRepository;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;
import org.springframework.stereotype.Service;

@Service
public class BookService {
  private BookRepository repository;

  public BookService(BookRepository repository) {
    this.repository = repository;
  }

  public List<Book> getAll() {
    return repository.findAll().stream()
        .map(b -> new Book(b.getId(), b.getTitle(), b.getAuthor(), b.getCategory()))
        .collect(Collectors.toList());
  }

  public Optional<Book> get(Long id) {
    return repository.findById(id)
        .map(b -> new Book(b.getId(), b.getTitle(), b.getAuthor(), b.getCategory()));
  }
}
