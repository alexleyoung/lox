#ifndef clox_chunk_h
#define clox_chunk_h

#include <stddef.h>
#include <stdint.h>

typedef enum {
  OP_RETURN,
} OpCode;

typedef struct {
  size_t count, capacity;
  uint8_t *code;
} Chunk;

void init_chunk(Chunk *chunk);
void write_chunk(Chunk *chunk, uint8_t byte);
void free_chunk(Chunk *chunk);

#endif
