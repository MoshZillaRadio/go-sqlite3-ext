cmake_minimum_required(VERSION 2.8)

project(go-sqlite3-ext)
add_library(go-sqlite3-ext SHARED src/fileio.c)
install(TARGETS go-sqlite3-ext DESTINATION lib)
