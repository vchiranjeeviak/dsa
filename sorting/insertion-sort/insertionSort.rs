fn insertion_sort(arr: &mut Vec<i32>) {
    for i in 1..arr.len() {
        let mut j = i - 1;
        let key = arr[i];
        // No need to check j >= 0 since j is of type usize and it can't be < 0
        while arr[j] > key {
            arr[j + 1] = arr[j];
            // checking j > 0  to avoid going negative since j is usize
            if j > 0 {
                j -= 1;
            } else {
                break;
            }
        }
        if j == 0 {
            arr[j] = key
        } else {
            arr[j + 1] = key;
        }
    }
}

fn main() {
    let mut arr = vec!(7, 7, 4, -2, 9);
    // Passing pointer instead of transferring ownership and getting it back
    insertion_sort(&mut arr);
    println!("{:?}", arr);
}