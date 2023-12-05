use std::time::{Instant};
use psutil::cpu;
use psutil::memory;

use std::thread;
use std::time::Duration;
use jemalloc_ctl::{stats, epoch};




fn main() {
    let start_time = Instant::now();
    run_your_code();
    let elapsed_time = Instant::now() - start_time;
    println!("Elapsed Time: {:?} milliseconds", elapsed_time.as_millis());
    fn_perf()
}

fn run_your_code() {
    // Allocate a large vector to consume memory
    let n = 1000 * 1000 * 1000;
    let mut large_vector = Vec::with_capacity(n);

    for _i in 0..n {
        large_vector.push(calculate_product());
    }
}

fn calculate_product() -> usize {
    let mut product = 1;
    for i in 1..=100 {
        product *= i;
    }
    product
}


pub fn fn_perf() {
    // Get CPU usage
    let cpu_percent = cpu::cpu_count();
    println!("CPU Usage: {}%", cpu_percent);

    loop {
        // many statistics are cached and only updated when the epoch is advanced.
        epoch::advance().unwrap();

        let allocated = stats::allocated::read().unwrap();
        let resident = stats::resident::read().unwrap();
        println!("{} bytes allocated/{} bytes resident", allocated, resident);
        thread::sleep(Duration::from_secs(10));
    }



    // Get memory usage
    // let memory_info = memory::virtual_memory().unwrap();
    // println!("Memory Usage: {} MB / {} MB", memory_info.used()/1024/1024, memory_info.total()/1024/1024);

}

