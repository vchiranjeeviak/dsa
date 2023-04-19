fn linear_search(arr: &Vec<i32>, target: i32) -> i32 {
    let mut count = 0;
    for i in arr {
        if *i == target {
            return count;
        }
        count += 1;
    }
    return -1;
}
fn main() {
    let arr = vec!(1, 7, 4, 2, 9);

    println!("{}", linear_search(&arr, 4))
}