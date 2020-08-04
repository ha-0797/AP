-- ========================================================================================================================== --


--
--                                                          ASSIGNMENT 1
--
--      A common type of text alignment in print media is "justification", where the spaces between words, are stretched or
--      compressed to align both the left and right ends of each line of text. In this problem we'll be implementing a text
--      justification function for a monospaced terminal output (i.e. fixed width font where every letter has the same width).
--
--      Alignment is achieved by inserting blanks and hyphenating the words. For example, given a text:
--
--              "He who controls the past controls the future. He who controls the present controls the past."
--
--      we want to be able to align it like this (to a width of say, 15 columns):
--
--              He who controls
--              the  past cont-
--              rols  the futu-
--              re. He  who co-
--              ntrols the pre-
--              sent   controls
--              the past.
--


-- ========================================================================================================================== --


import Data.List
import Data.Char
import Data.Maybe

text1 = "He who controls the past controls the future. He who controls the present controls the past."
text2 = "A creative man is motivated by the desire to achieve, not by the desire to beat others."


-- ========================================================================================================================== --







-- ========================================================= PART 1 ========================================================= --


--
-- Define a function that splits a list of words into two lists, such that the first list does not exceed a given line width.
-- The function should take an integer and a list of words as input, and return a pair of lists.
-- Make sure that spaces between words are counted in the line width.
--
-- Example:
--    splitLine ["A", "creative", "man"] 12   ==>   (["A", "creative"], ["man"])
--    splitLine ["A", "creative", "man"] 11   ==>   (["A", "creative"], ["man"])
--    splitLine ["A", "creative", "man"] 10   ==>   (["A", "creative"], ["man"])
--    splitLine ["A", "creative", "man"] 9    ==>   (["A"], ["creative", "man"])
--
ss :: [String] -> Int -> [String]
ss x y = snd(splitLine x y)

fs :: [String] -> Int -> [String]
fs x y = fst(splitLine x y)

n_length:: String -> Int -> Int
n_length x y = y - length x - 1

splitLine :: [String] -> Int -> ([String], [String])
-- Function definition here

splitLine [] _ = ([],[])
splitLine (x:xs) y  
  | length x <= y  = ([x] ++ fs xs (n_length x y), ss xs (n_length x y))
  | otherwise = (fs xs 0, [x] ++ ss xs 0)
-- ========================================================= PART 2 ========================================================= --


--
-- To be able to align the lines nicely. we have to be able to hyphenate long words. Although there are rules for hyphenation
-- for each language, we will take a simpler approach here and assume that there is a list of words and their proper hyphenation.
-- For example:

--
-- Define a function that splits a list of words into two lists in different ways. The first list should not exceed a given
-- line width, and may include a hyphenated part of a word at the end. You can use the splitLine function and then attempt
-- to breakup the next word with a given list of hyphenation rules. Include a breakup option in the output only if the line
-- width constraint is satisfied.
-- The function should take a hyphenation map, an integer line width and a list of words as input. Return pairs of lists as
-- in part 1.
--
-- Example:
--    		10	 ==>   [(["He","who"], ["controls."]), (["He","who","co-"], ["ntrols."]), (["He","who","cont-"], ["rols."])]
--
-- Make sure that words from the list are hyphenated even when they have a trailing punctuation (e.g. "controls.")
--
-- You might find 'map', 'find', 'isAlpha' and 'filter' useful.
--
enHyp = [("creative", ["cr","ea","ti","ve"]), ("controls", ["co","nt","ro","ls"]), ("achieve", ["ach","ie","ve"]), ("future", ["fu","tu","re"]), ("present", ["pre","se","nt"]), ("motivated", ["mot","iv","at","ed"]), ("desire", ["de","si","re"]), ("others", ["ot","he","rs"])]

search :: String -> [(String,[String])] -> [String]
search _ [] = []
search a (x:xs) 
  | a == fst(x) = snd(x)
  | otherwise = search a xs

con :: [String] -> Int -> String
con [] _ = []
con (x:xs) y 
  | length (x) < y = x ++ con xs (y - length x)
  | otherwise = "-"

list_len :: [String] -> Int
list_len [] = 0
list_len (x:xs) = length x + list_len xs + 1

con_len :: [String] -> Int -> Int -> Int
con_len [] y _ = y
con_len (x:xs) y z
  | length x < y = con_len xs (y - length x) (0 + length x)
  | otherwise = z

give_first :: [a] -> a
give_first (x:xs) = x

rem_punct :: String -> String
rem_punct [] = []
rem_punct (x:xs) 
  | isAlpha(x) = [x] ++ rem_punct xs
  | otherwise = "" 

search' :: [String] ->  Int -> String
search' x y = rem_punct(give_first(ss x y))

con2 x y = con (search (search' x y) enHyp) (y - list_len (fs x y))

del :: String -> String -> String
del y xs = [x 
  | x <- xs, not (x `elem` y)]

one :: [String] -> [String]
one (x:xs) = [x]
one [] = []

two :: [String] -> [String]
two (x:xs) = xs
two [] = []

del' :: String -> [String] -> [String]
del' y (x:xs) = [del y x] 

stuff :: [String] -> Int -> [([String],[String])]
stuff x y = [(fs x y ++ [con2 x y], del' (rem_punct(con2 x y)) (one (ss x y)) ++ two (ss x y))]

check :: Int -> [String] -> Int
check y x = y - (list_len (fs x y)) - con_len (search (search' x y) enHyp) (y - list_len (fs x y) + 1) y

check2 :: [String] -> Int -> Int
check2 x y = y - con_len (search (search' x y) enHyp) (y - list_len (fs x y) + 1) 0

lineBreaks :: [(String, [String])] -> Int -> [String] -> [([String], [String])]
-- Function definition here
lineBreaks enHyp y [] = [([],[])]
lineBreaks enHyp y x
  | check y x > 0 = lineBreaks enHyp (check2 x y) x ++ stuff x y
  | otherwise = [splitLine x y] 
-- ========================================================= PART 3 ========================================================= --


--
-- Define a function that inserts a given number of blanks (spaces) into a list of strings and outputs a list of all possible
-- insertions. Only insert blanks between strings and not at the beginning or end of the list (if there are less than two
-- strings in the list then return nothing). Remove duplicate lists from the output.
-- The function should take the number of blanks and the the list of strings as input and return a lists of strings.
--
-- Example:
--    blankInsertions 2 ["A", "creative", "man"]   ==>   [["A", " ", " ", "creative", "man"], ["A", " ", "creative", " ", "man"], ["A", "creative", " ", " ", "man"]]
--
-- Use let/in/where to make the code readable
--

removeDuplicates2 = foldl (\seen x -> if x `elem` seen
                                      then seen
                                      else seen ++ [x]) []

addToAll :: String -> [[String]] -> [[String]]
addToAll y (x:xs) = [[y]++x] ++ addToAll y xs
addToAll _ [] = []

insertOne :: [String] -> [[String]]
insertOne (x:y:xs) = [[x] ++ [" "] ++ [y] ++ xs] ++ addToAll x (insertOne (y:xs))
insertOne (x:xs) = []
insertOne [] = [[]]

insertToList :: [[String]] -> [[String]]
insertToList (x:xs) = insertOne x ++ insertToList xs
insertToList [] = [] 

blankInsertions :: Int -> [String] -> [[String]]
-- Function definition here
blankInsertions _ [] = []
blankInsertions 0 x = [x]
blankInsertions y x 
  | y == 1 = insertOne x
  | otherwise = removeDuplicates2 (insertToList (blankInsertions (y - 1) x))




-- ========================================================= PART 4 ========================================================= --


--
-- Define a function to score a list of strings based on four factors:
--
--    blankCost: The cost of introducing each blank in the list
--    blankProxCost: The cost of having blanks close to each other
--    blankUnevenCost: The cost of having blanks spread unevenly
--    hypCost: The cost of hyphenating the last word in the list
--
-- The cost of a list of strings is computed simply as the weighted sum of the individual costs. The blankProxCost weight equals
-- the length of the list minus the average distance between blanks (0 if there are no blanks). The blankUnevenCost weight is
-- the variance of the distances between blanks.
--
-- The function should take a list of strings and return the line cost as a double
--
-- Example:
--    lineCost ["He", " ", " ", "who", "controls"]
--        ==>   blankCost * 2.0 + blankProxCost * (5 - average(1, 0, 2)) + blankUnevenCost * variance(1, 0, 2) + hypCost * 0.0
--        ==>   blankCost * 2.0 + blankProxCost * 4.0 + blankUnevenCost * 0.666...
--
-- Use let/in/where to make the code readable
--


---- Do not modify these in the submission ----
blankCost = 1.0
blankProxCost = 1.0
blankUnevenCost = 1.0
hypCost = 1.0
-----------------------------------------------

noBlank :: [String] -> Double
noBlank [] = 0.0
noBlank (x:xs)
  | x == " " = 1.0 + noBlank xs
  | otherwise = noBlank xs

dist :: [String] -> Double -> [Double]
dist [] y = [y]
dist (x:xs) y
  | x == " " = [y] ++ dist xs 0.0
  | otherwise = dist xs (y + 1.0)

avg :: [Double] -> Double
avg [] = 0.0 
avg (x:xs) = x + avg xs

varience :: [Double] -> Double -> Double
varience [] y = y
varience (x:xs) y
  | x > y = varience xs x
  | otherwise = varience xs y

isHyp' :: String -> Bool
isHyp' [] = False
isHyp' (x:xs)
  | x == '-' = True
  | otherwise = isHyp' xs

isHyp :: [String] -> Double
isHyp [] = 0.0
isHyp (x:xs)
  | isHyp' x = 1.0
  |otherwise = 0.0 + isHyp xs


lineCost :: [String] -> Double
-- Function definition here
lineCost [] = 0.0
lineCost x = blankCost * (noBlank x) + blankProxCost * (fromIntegral(length x) - (avg (dist(x) 0))/ (fromIntegral(length (dist x 0)))) + blankUnevenCost * (varience (dist x 0) 0/ (fromIntegral(length (dist x 0)))) + hypCost * (isHyp x)





-- ========================================================= PART 5 ========================================================= --


--
-- Define a function that returns the best line break in a list of words given a cost function, a hyphenation map and the maximum
-- line width (the best line break is the one that minimizes the line cost of the broken list).
-- The function should take a cost function, a hyphenation map, the maximum line width and the list of strings to split and return
-- a pair of lists of strings as in part 1.
--
-- Example:
--    bestLineBreak lineCost enHyp 12 ["He", "who", "controls"]   ==>   (["He", "who", "cont-"], ["rols"])
--
-- Use let/in/where to make the code readable

--bestLineBreak :: ([String] -> Double) -> [(String, [String])] -> Int -> [String] -> ([String], [String])
-- Function definition here
--bestLineBreak lineCost enHyp y x = choose (lineBreaks enHyp y x) (fst (lineBreaks enHyp y x))

-- Finally define a function that justifies a given text into a list of lines satisfying a given width constraint.
-- The function should take a cost function, hyphenation map, maximum line width, and a text string as input and return a list of
-- strings.
--
-- 'justifyText lineCost enHyp 15 text1' should give you the example at the start of the assignment.
--
-- You might find the words and unwords functions useful.
--


--justifyText :: ([String] -> Double) -> [(String, [String])] -> Int -> String -> [String]
-- Function definition here
