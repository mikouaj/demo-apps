package dev.stefaniak.demo.service;

public class Book {
  private Long id;
  private String title;
  private String author;
  private String category;

  public Book(Long id, String title, String author, String category) {
    this.id = id;
    this.title = title;
    this.author = author;
    this.category = category;
  }

  public Long getId() {
    return id;
  }

  public void setId(Long id) {
    this.id = id;
  }

  public String getTitle() {
    return title;
  }

  public void setTitle(String title) {
    this.title = title;
  }

  public String getAuthor() {
    return author;
  }

  public void setAuthor(String author) {
    this.author = author;
  }

  public String getCategory() {
    return category;
  }

  public void setCategory(String category) {
    this.category = category;
  }
}
