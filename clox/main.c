#include "chunk.h"

int main() {
  Chunk chunk;
  init_chunk(&chunk);
  write_chunk(&chunk, OP_RETURN);
  free_chunk(&chunk);
}
