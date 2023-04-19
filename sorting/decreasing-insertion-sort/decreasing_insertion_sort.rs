fn decreasing_insertion_sort(arr: &mut Vec<i32>) {
    for i in 1..arr.len() {
        let mut j = i - 1;
        let key = arr[i];
        // j >= 0 throws warning since j is of type usize and it can't be < 0
        while j > 0 && arr[j] < key {
            arr[j + 1] = arr[j];
            j -= 1;
        }
        arr[j + 1] = key;
        // Checking just j = 0 case
        if j == 0 && arr[j] < key {
            arr[1] = arr[0];
            arr[0] = key;
        }
    }
}

fn main() {
    let mut arr = vec![9, 5, -1, 7, 0];
    decreasing_insertion_sort(&mut arr);
    println!("{:?}", arr);
}
