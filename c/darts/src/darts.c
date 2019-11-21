#include "darts.h"


uint8_t score(struct coordinate_t landing_position) {

    float position = (landing_position.x *landing_position.x) + (landing_position.y*landing_position.y);

    if (position <= 1) {
        return 10;
    } 
    else if (position <= 25) {
        return 5;
    }
    else if (position <= 100) {
        return 1;
    }
    else {
        return 0;
    }
}
