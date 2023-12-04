use std::time::{Instant};
// use psutil::cpu;
// use psutil::memory;


fn main() {
    let start_time = Instant::now();
    run_your_code();
    let elapsed_time = Instant::now() - start_time;
    println!("Elapsed Time: {:?} milliseconds", elapsed_time.as_millis());
}

fn run_your_code() {
    // Allocate a large vector to consume memory
    let n = 1000 * 1000 * 1000;
    println!("Array Size: {}", n);
    let mut large_vector = Vec::with_capacity(n);

    for i in 0..n {
        large_vector.push(i);
    }

    // CPU-intensive operation - calculating Pi using the Leibniz formula
    let mut pi : f64 = 0.0;
    for i in 0..n {
        pi += if i % 2 == 0 { 1.0 } else { -1.0 } * 4.0 / (2 * i + 1) as f64;
    }
}


// pub fn fn_perf() {
//     // Get CPU usage
//     let cpu_percent = cpu::cpu_count();
//     println!("CPU Usage: {}%", cpu_percent);

//     // Get memory usage
//     let memory_info = memory::virtual_memory().unwrap();
//     println!("Memory Usage: {} B / {} B", memory_info.used() >> 20, memory_info.total() >> 20);

// }

