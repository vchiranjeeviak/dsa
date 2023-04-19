use std::mem;

fn reverse(mut arr: Vec<i8>) -> Vec<i8> {
    let len = arr.len();
    for i in 0..len / 2 {
        let temp = arr[i];
        arr[i] = arr[len - i - 1];
        arr[len - i - 1] = temp;
    }
    return arr;
}

fn binary_addition_arrays(mut arr1: Vec<i8>, mut arr2: Vec<i8>) -> Vec<i8> {
    let mut len1 = arr1.len();
    let mut len2 = arr2.len();

    if len1 > len2 {
        mem::swap(&mut len1, &mut len2);
        mem::swap(&mut arr1, &mut arr2);
    }

    let mut res = vec![];
    let mut carry = 0;

    for i in 0..len2 {
        let mut sum = arr2[len2 - i - 1] + carry;
        if len1 - i > 0 {
            sum += arr1[len1 - i - 1];
        }
        if sum > 1 {
            carry = 1;
            sum %= 2;
        }
        res.push(sum);
    }
    res.push(carry);
    res = reverse(res);
    return res;
}

fn main() {
    let arr1 = vec![1, 0, 1];
    let arr2 = vec![1, 1, 1];
    println!("{:?}", binary_addition_arrays(arr1, arr2));
}
