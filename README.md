# Project Infininty

## Summary
Small project I made to practice creating and deploying a fullstack application to AWS

- Frontend in Nuxt.js and Vue
- Backend in Go

### Notes for self
- To make the docker image appear in list run the following:
  - `docker build -t go_build . --load`
- To make the docker image use the intended port:
  - `docker run -dp 8080:8080 go_build`