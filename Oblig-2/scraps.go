
/* socket loop logo
Tutorials
Golang : How to count duplicate items in slice/array?

5602 views  2nd June 2015

Tags : golang count-duplicate slice array 

Problem :

You need to count the number of duplicate items in a slice or array.

Solution :

Pseudo-code : Create a map and insert one item from the slice/array with a for loop. Before inserting a new item check if a similar item already exist in the map. If yes, increase the counter value associated with the particular item and if no, assign a new counter with the value of 1 for the new item(in the map).

Here you go :
*/ 

 package main

 import (
  "fmt"
 )

 func printslice(slice []string) {
  fmt.Println("slice = ", slice)
 }

 func dup_count(list []string) map[string]int {

  duplicate_frequency := make(map[string]int)

  for _, item := range list {
 // check if the item/element exist in the duplicate_frequency map

 _, exist := duplicate_frequency[item]

 if exist {
 duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
 } else {
 duplicate_frequency[item] = 1 // else start counting from 1
 }
  }
  return duplicate_frequency
 }

 func main() {
  duplicate := []string{"Hello", "World", "GoodBye", "World", "We", "Love", "Love", "You"}

  printslice(duplicate)

  dup_map := dup_count(duplicate)

  //fmt.Println(dup_map)

  for k, v := range dup_map {
 fmt.Printf("Item : %s , Count : %d\n", k, v)
  }

 }
Sample output :

slice = [Hello World GoodBye World We Love Love You]

Item : You , Count : 1

Item : Hello , Count : 1

Item : World , Count : 2

Item : GoodBye , Count : 1

Item : We , Count : 1

Item : Love , Count : 2

