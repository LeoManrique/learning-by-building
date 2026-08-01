# 03 — Normalization

**Goal:** one helper both checks share: lowercase, letters only. This is the core of the project.

- First internalize what a string is in your language: iterate over `"Aé!b"` character by character, logging each character and its position, and notice the positions are byte offsets that can skip — that's UTF-8 encoding, and it's why you work in characters, not bytes.
- Build the normal form from the input: keep only letters, lowercased. Your standard library has a way to test whether a character is a letter and a way to lowercase one — find them.
- Log input → normal form for the panama string and check you get `amanaplanacanalpanama`. Try a few of your own, including one with digits and punctuation.

**Done when:**

- [ ] The helper turns `"A man a plan a canal Panama"` into `amanaplanacanalpanama`. No new tests pass yet — both checks will stand on this.
