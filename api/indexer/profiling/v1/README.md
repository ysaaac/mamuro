# Indexing V1

For this first approach it gets all headers of the email such as:

```zsh
"Message-ID"
"Date"
"From"
"To"
"Subject"
"Mime-Version"
"Content-Type"
"Content-Transfer-Encoding"
"X-From"
"X-To"
"X-cc"
"X-bcc"
"X-Folder"
"X-Origin"
"X-FileName"
```

and also consider all behind this as message content.

- Consideration:
    - I'm using `bulkV2` to populate the indices instead of `create` for performance purposes when sending data
      to `Zincsearch`