# Learning by Building

A custom, project-based path from CLI tools to advanced stuff.

## Motivation

Learning about deeper concepts that are usually taking for granted when building software.

## Scope of projects

Projects that help build foundations and apply important concepts. If some projects feel irrelevant they will be deleted from curriculum.

## Pace

Probably years. I don't expect to complete all the projects.

## Development

### Tools

**Programming language:** Go & Rust.
**AI Assistant:** Claude Code.

### AI Usage

***AI will be used for:**
- Explain concepts
- Help with syntax
- Suggest best practices
- Mark projects as completed based on acceptance criteria
- Write repetitive code
- Write, run and review outputs of E2E tests

***AI will NOT be used for:**
- Write logic
- Write code as an agent (every edit must be reviewed)

## Repo structure

Monorepo. Every single project will be built in a subfolder of the main repo until they matched the minimal acceptance criteria. If a project is later used more seriously then it will have an enhanced version in a separate repo, but will not be part of the Learning by Building project.

## Plan

- [curriculum.yaml](curriculum.yaml) — the full curriculum: every project with description, concepts, acceptance criteria, and recommended languages. Designed by Claude, Gemini and ChatGPT in Apr-2026.

## How to read a project entry

Each project has:
- `description` — what to build, in plain words
- `concepts` — the ideas the project drives into your head
- `requirements` — concrete acceptance criteria; tick them all and you're done
- `recommended_languages` — good fits, in rough order of preference

## Target audience

Myself.