package dev.stefaniak.demo.rest;

import dev.stefaniak.demo.service.BookService;
import java.util.List;
import java.util.stream.Collectors;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class BookRestController {
  private BookService service;

  @ResponseStatus(value = HttpStatus.NOT_FOUND)
  public class ResourceNotFoundException extends RuntimeException {
  }
  public BookRestController(BookService service) {
    this.service = service;
  }

  @GetMapping("/books")
  public List<Book> books() {
    return service.getAll().stream()
        .map(b -> new Book(b.getId(), b.getTitle(), b.getAuthor(), b.getCategory()))
        .collect(Collectors.toList());
  }

  @GetMapping("/book/{id}")
  public Book book(@PathVariable Long id) {
    return service.get(id)
        .map(b -> new Book(b.getId(), b.getTitle(), b.getAuthor(), b.getCategory()))
        .orElseThrow(ResourceNotFoundException::new);
  }
}
