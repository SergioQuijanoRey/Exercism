// Package listops implements list operations
package listops


// * `append` (*given two lists, add all items in the second list to the end of the first list*);
// * `concatenate` (*given a series of lists, combine all items in all lists into one flattened list*);
// * `filter` (*given a predicate and a list, return the list of all items for which `predicate(item)` is True*);
// * `length` (*given a list, return the total number of items within it*);
// * `map` (*given a function and a list, return the list of the results of applying `function(item)` on all items*);
// * `foldl` (*given a function, a list, and initial accumulator, fold (reduce) each item into the accumulator from the left using `function(accumulator, item)`*);
// * `foldr` (*given a function, a list, and an initial accumulator, fold (reduce) each item into the accumulator from the right using `function(item, accumulator)`*);
// * `reverse` (*given a list, return a list with all the original items, but in reversed order*);

func append(first *[]interface{}, second *[]interface{}){
    for _, val := range(second){
        first.
    }
}
