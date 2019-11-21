#ifndef DARTS_H
#define DARTS_H

#include <stdint.h>

typedef struct coordinate_t coordinate_t;

struct coordinate_t
{
    float x, y;
};

uint8_t score(struct coordinate_t landing_position);

#endif
