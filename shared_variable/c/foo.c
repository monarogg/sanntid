// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t mu = PTHREAD_MUTEX_INITIALIZER;

// Note the return type: void*
void* incrementingThreadFunction(){
    for (int j = 0; j < 1000000; j++) {
        pthread_mutex_lock(&mu);
        i++;
        pthread_mutex_unlock(&mu);
    }
    return NULL;
}

void* decrementingThreadFunction(){
    for (int j = 0; j < 1000000; j++) {
        pthread_mutex_lock(&mu);
        i--;
        pthread_mutex_unlock(&mu);
    }
    return NULL;
}


int main(){
    // TODO: 
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?
    pthread_t thread1, thread2;

    // thread1 for incrementing:
    if (pthread_create(&thread1, NULL, incrementingThreadFunction, NULL)) {
        fprintf(stderr, "Error creating thread1 1\n");
        return 1;
    }
    
    // thread2 for decrementing:
    if (pthread_create(&thread2, NULL, decrementingThreadFunction, NULL)) {
        fprintf(stderr, "Error creating thread2 1\n");
        return 1;
    }


    // vente på at begge trådene blir ferdige: 
    if (pthread_join(thread1, NULL)) {
        fprintf(stderr, "Error joining thread 1\n");
        return 1;
    }

    if (pthread_join(thread2, NULL)) {
        fprintf(stderr, "Error joining thread 1\n");
        return 1;
    }
    
    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`    
    
    printf("The magic number is: %d\n", i);

    pthread_mutex_destroy(&mu);

    return 0;
}
