use std::time::{Instant};
use psutil::cpu;
use psutil::memory;


fn main() {
    elapsed_time();
    // fn_perf();

}

fn elapsed_time() {
    let start_time = Instant::now();

    // Allocate memory
    let array_size = 1_000_000_000;
    let mut memory_array = Vec::with_capacity(array_size);

    for i in 1..array_size {
        memory_array.push(i);
    }

    let elapsed_time = (Instant::now() - start_time)*1000;

    println!("Array Size: {}", array_size);
    println!("Elapsed Time: {:?}", elapsed_time);
}

pub fn fn_perf() {
    // Get CPU usage
    let cpu_percent = cpu::cpu_count();
    println!("CPU Usage: {}%", cpu_percent);

    // Get memory usage
    let memory_info = memory::virtual_memory().unwrap();
    println!("Memory Usage: {} B / {} B", memory_info.used() >> 20, memory_info.total() >> 20);

}

