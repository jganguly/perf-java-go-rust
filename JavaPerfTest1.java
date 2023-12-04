import java.lang.management.ManagementFactory;
import java.lang.management.MemoryMXBean;
import java.lang.management.MemoryUsage;
import java.lang.management.OperatingSystemMXBean;

public class JavaPerfTest1 {

    public static void main(String[] args) {
        elapsedTime();

        // Get the OperatingSystemMXBean
        // OperatingSystemMXBean osBean = ManagementFactory.getOperatingSystemMXBean();
        // // Print CPU Usage
        // System.out.println("CPU Usage: " + osBean.getSystemLoadAverage() + "%");

        // // Get the MemoryMXBean
        // MemoryMXBean memoryBean = ManagementFactory.getMemoryMXBean();
        // // Print Heap Memory Usage
        // MemoryUsage heapMemoryUsage = memoryBean.getHeapMemoryUsage();
        // System.out.println("Heap Memory Usage: " + formatMemoryUsage(heapMemoryUsage));

        // Print Non-Heap Memory Usage
        // MemoryUsage nonHeapMemoryUsage = memoryBean.getNonHeapMemoryUsage();
        // System.out.println("Non-Heap Memory Usage: " + formatMemoryUsage(nonHeapMemoryUsage));
    }


    private static String formatMemoryUsage(MemoryUsage memoryUsage) {
        long max = memoryUsage.getMax() / (1024 * 1024);
        long used = memoryUsage.getUsed() / (1024 * 1024);
        long committed = memoryUsage.getCommitted() / (1024 * 1024);

        return "Used: " + used + "MB, Committed: " + committed + "MB, Max: " + max + "MB";
    }


    private static void elapsedTime() {
        long startTime = System.currentTimeMillis();

        // Allocate memory
        int arraySize = 1000*1000*1000;
        int[] memoryArray = new int[arraySize];

        for (int i=0; i<arraySize; i++) {
            memoryArray[i] = i;
        }

        long endTime = System.currentTimeMillis();
        long elapsedTime = endTime - startTime;

        System.out.println("Output: Java21");
        System.out.println("Array Size: " + arraySize);
        System.out.println("Elapsed Time: " + elapsedTime + " milliseconds");
    }
}
