#include "chunk.h"
#include "debug.h"

int main() {
  Chunk chunk;
  init_chunk(&chunk);
  write_chunk(&chunk, OP_RETURN);
  disassemble_chunk(&chunk, "test");
  free_chunk(&chunk);
}
