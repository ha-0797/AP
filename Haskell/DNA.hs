-- ---------------------------------------------------------------------
-- DNA Analysis 
-- CS300 Spring 2018
-- Due: 24 Feb 2018 @9pm
-- ------------------------------------Assignment 2------------------------------------
--
-- >>> YOU ARE NOT ALLOWED TO IMPORT ANY LIBRARY
-- Functions available without import are okay
-- Making new helper functions is okay
--
-- ---------------------------------------------------------------------
--
-- DNA can be thought of as a sequence of nucleotides. Each nucleotide is 
-- adenine, cytosine, guanine, or thymine. These are abbreviated as A, C, 
-- G, and T.
--
type DNA = [Char]
type RNA = [Char]
type Codon = [Char]
type AminoAcid = Maybe String

-- ------------------------------------------------------------------------
-- 				PART 1
-- ------------------------------------------------------------------------				

-- We want to calculate how alike are two DNA strands. We will try to 
-- align the DNA strands. An aligned nucleotide gets 3 points, a misaligned
-- gets 2 points, and inserting a gap in one of the strands gets 1 point. 
-- Since we are not sure how the next two characters should be aligned, we
-- try three approaches and pick the one that gets us the maximum score.
-- 1) Align or misalign the next nucleotide from both strands
-- 2) Align the next nucleotide from first strand with a gap in the second     
-- 3) Align the next nucleotide from second strand with a gap in the first    
-- In all three cases, we calculate score of leftover strands from recursive 
-- call and add the appropriate penalty.                                    

match = 3
mismatch = 2
gap = 1;

score :: DNA -> DNA -> Int
score (x:xs) (y:ys)
  | y==x = maximum[(score xs ys + match), (score (x:xs) ys + gap), (score xs (y:ys)+gap)]
  | otherwise = maximum[(score xs ys + mismatch), (score (x:xs) ys + gap), (score xs (y:ys)+gap)]
score [] (y:ys) = score [] ys + gap
score (x:xs) [] = score xs [] + gap
score [] [] = 0
-- -------------------------------------------------------------------------
--				PART 2
-- -------------------------------------------------------------------------
-- Write a function that takes a list of DNA strands and returns a DNA tree. 
-- For each DNA strand, make a separate node with zero score 
-- in the Int field. Then keep merging the trees. The merging policy is:
-- 	1) Merge two trees with highest score. Merging is done by making new
--	node with the smaller DNA (min), and the two trees as subtrees of this
--	tree
--	2) Goto step 1 :)
--

data DNATree = Node DNA Int DNATree DNATree | Nil deriving (Ord, Show, Eq)

--makeDNATree :: [DNA] -> DNATree

--test = ["AACCTTGG","ACTGCATG", "ACTACACC", "ATATTATA"]

-- -------------------------------------------------------------------------
--				PART 3
-- -------------------------------------------------------------------------

-- Even you would have realized it is hard to debug and figure out the tree
-- in the form in which it currently is displayed. Lets try to neatly print 
-- the DNATree. Each internal node should show the 
-- match score while leaves should show the DNA strand. In case the DNA strand 
-- is more than 10 characters, show only the first seven followed by "..." 
-- The tree should show like this for an evolution tree of
-- ["AACCTTGG","ACTGCATG", "ACTACACC", "ATATTATA"]
--
-- 
-- Make helper functions as needed. It is a bit tricky to get it right. One
-- hint is to pass two extra string, one showing what to prepend to next 
-- level e.g. "+---" and another to prepend to level further deep e.g. "|   "

level = "|   "
next = "+---"
draw :: DNATree -> [Char]
draw (Node str sc left right) = drawHelper (Node str sc left right) "" next 

drawHelper :: DNATree -> String -> String -> String
drawHelper (Node str 0 Nil Nil) x y = x ++ y ++ str ++ ['\n']
drawHelper (Node str sc left right) x y = x ++ y ++ (show sc) ++ ['\n'] ++ drawHelper left (x ++ level) y ++ drawHelper right (x ++ level) y

--putstrln 

-- ---------------------------------------------------------------------------
--				PART 4
-- ---------------------------------------------------------------------------
--
--
-- Our score function is inefficient due to repeated calls for the same 
-- suffixes. Lets make a dictionary to remember previous results. First you
-- will consider the dictionary as a list of tuples and write a lookup
-- function. Return Nothing if the element is not found. Also write the 
-- insert function. You can assume that the key is not already there.
type Dict a b = [(a,b)]

lookupDict :: (Eq a) => a -> Dict a b -> Maybe b
lookupDict y (x:xs) 
  | y == fst(x) = Just (snd(x))
  | otherwise = lookupDict y xs
lookupDict _ [] = Nothing 

insertDict :: (Eq a) => a -> b -> (Dict a b)-> (Dict a b)
insertDict x y z = z ++ [(x,y)]

-- We will improve the score function to also return the alignment along
-- with the score. The aligned DNA strands will have gaps inserted. You
-- can represent a gap with "-". You will need multiple let expressions 
-- to destructure the tuples returned by recursive calls.


maximum' :: [((String, String), Int)] -> ((String, String), Int)
maximum' (x:xs:xss) 
  | (snd(x)) > (snd(xs)) = maximum' (x : xss)
  | otherwise = maximum' (xs:xss)
maximum' (x:xs) = x

alignment :: String -> String -> ((String, String), Int)
alignment [] [] = (("",""),0)
alignment [] (y:ys) =
  let sc = snd(alignment [] ys)
      str1 = "-" ++ fst(fst(alignment [] ys)) 
      str2 = [y] ++ snd(fst(alignment [] ys))
  in  ((str1, str2), sc)
alignment (x:xs) [] =
  let sc = snd(alignment xs [])
      str1 = [x] ++ fst(fst(alignment xs [])) 
      str2 = "-" ++ snd(fst(alignment xs []))
  in  ((str1, str2), sc)

alignment (x:xs) (y:ys) 
  | x == y = 
  	let one = (( [x] ++ fst(fst(alignment xs ys)), [y] ++ snd(fst(alignment xs ys))), snd(alignment xs ys) + match)
  	    two = (( "-" ++ fst(fst(alignment (x:xs) ys)), [y] ++ snd(fst(alignment (x:xs) ys))), snd(alignment (x:xs) ys) + gap) 	--Sorry :p
  	    three = (([x]++fst(fst(alignment xs (y:ys))),"-" ++ snd(fst(alignment xs (y:ys)))), snd(alignment xs (y:ys)) + gap)
  	in  maximum' [one, two, three]
  |otherwise = 	
  	let one = (( [x] ++ fst(fst(alignment xs ys)), [y] ++ snd(fst(alignment xs ys))), snd(alignment xs ys) + mismatch)
  	    two = (( "-" ++ fst(fst(alignment (x:xs) ys)), [y] ++ snd(fst(alignment (x:xs) ys))), snd(alignment (x:xs) ys) + gap) 	
  	    three = (([x]++fst(fst(alignment xs (y:ys))),"-" ++ snd(fst(alignment xs (y:ys)))), snd(alignment xs (y:ys)) + gap)
  	in  maximum' [one, two, three]


-- We will now pass a dictionary to remember previously calculated scores 
-- and return the updated dictionary along with the result. Use let 
-- expressions like the last part and pass the dictionary from each call
-- to the next. Also write logic to skip the entire calculation if the 
-- score is found in the dictionary. You need just one call to insert.
type ScoreDict = Dict (DNA,DNA) Int

findList :: (DNA, DNA) -> ScoreDict -> Bool
findList y (x:xs)
  | ((fst(y)) == (fst(fst(x))) && (snd(y)) == (snd(fst(x)))) = True
  | otherwise = findList y xs
findList _ [] = False

scoreMemo :: (DNA,DNA) -> ScoreDict -> (ScoreDict,Int)
scoreMemo x y
  | findList x y = (y, (score (fst x) (snd x)))
  | otherwise = ((insertDict x (score (fst x) (snd x)) y), (score (fst x) (snd x)))


-- In this part, we will use an alternate representation for the 
-- dictionary and rewrite the scoreMemo function using this new format.
-- The dictionary will be just the lookup function so the dictionary 
-- can be invoked as a function to lookup an element. To insert an
-- element you return a new function that checks for the inserted
-- element and returns the old dictionary otherwise. You will have to
-- think a bit on how this will work. An empty dictionary in this 
-- format is (\_->Nothing)


type Dict2 a b = a->Maybe b


isNothing :: Maybe a -> Bool
isNothing (Just x) = False
isNothing Nothing = True


insertDict2 :: (Eq a) => a -> b -> (Dict2 a b)-> (Dict2 a b)
insertDict2 x y f
  | isNothing (f(x)) = 
  	let new_f(x) = Just y
  	in  new_f 
  | otherwise = f

type ScoreDict2 = Dict2 (DNA,DNA) Int

scoreMemo2 :: (DNA,DNA) -> ScoreDict2 -> (ScoreDict2,Int)
scoreMemo2 x y
  | (y(x)) == (Just (score (fst(x)) (snd(x)))) = (y, (score (fst x) (snd x)))
  | otherwise = (insertDict2 x (score (fst x) (snd x)) y, (score (fst x) (snd x)))

-- ---------------------------------------------------------------------------
-- 				PART 5
-- ---------------------------------------------------------------------------

-- Now, we will try to find the mutationDistance between two DNA sequences.
-- You have to calculate the number of mutations it takes to convert one 
-- (start sequence) to (end sequence). You will also be given a bank of 
-- sequences. However, there are a couple of constraints, these are as follows:

-- 1) The DNA sequences are of length 8
-- 2) For a sequence to be a part of the mutation distance, it must contain 
-- "all but one" of the neuclotide bases as its preceding sequence in the same 
-- order AND be present in the bank of valid sequences
-- 'AATTGGCC' -> 'AATTGGCA' is valid only if 'AATTGGCA' is present in the bank
-- 3) Assume that the bank will contain valid sequences and the start sequence
-- may or may not be a part of the bank.
-- 4) Return -1 if a mutation is not possible

	
-- mutationDistance "AATTGGCC" "TTTTGGCA" ["AATTGGAC", "TTTTGGCA", "AAATGGCC" "TATTGGCC", "TTTTGGCC"] == 3
-- mutationDistance "AAAAAAAA" "AAAAAATT" ["AAAAAAAA", "AAAAAAAT", "AAAAAATT", "AAAAATTT"] == 2

findDistance :: DNA -> DNA -> Int
findDistance (x:xs) (y:ys)
  | x == y = findDistance xs ys
  | otherwise = 1 + findDistance xs ys
findDistance [] [] = 0

mutationDistance :: DNA -> DNA -> [DNA] -> Int
mutationDistance a b (x:xs)
  | b == x = findDistance a b
  | otherwise = mutationDistance a b xs
mutationDistance a b [] = -1


-- ---------------------------------------------------------------------------
-- 				PART 6
-- ---------------------------------------------------------------------------
--
-- Now, we will write a function to transcribe DNA to RNA. 
-- The difference between DNA and RNA is of just one base i.e.
-- instead of Thymine it contains Uracil. (U)
--
transcribeDNA :: DNA -> RNA
transcribeDNA [] = ""
transcribeDNA (x:xs) 
  | x == 'T' = "U" ++ transcribeDNA xs
  | otherwise = [x] ++ transcribeDNA xs

-- Next, we will translate RNA into proteins. A codon is a group of 3 neuclotides 
-- and forms an aminoacid. A protein is made up of various amino acids bonded 
-- together. Translation starts at a START codon and ends at a STOP codon. The most
-- common start codon is AUG and the three STOP codons are UAA, UAG and UGA.
-- makeAminoAcid should return Nothing in case of a STOP codon.
-- Your translateRNA function should return a list of proteins present in the input
-- sequence. 
-- Please note that the return type of translateRNA is [String], you should convert
-- the abstract type into a concrete one.
-- You might wanna use the RNA codon table from 
-- https://www.news-medical.net/life-sciences/RNA-Codons-and-DNA-Codons.aspx
-- 
--
makeAminoAcid :: Codon -> AminoAcid
makeAminoAcid x 
  | x == "UUU" = Just "Phe"
  | x == "UUC" = Just "Phe"
  | x == "UUA" = Just "Leu"
  | x == "UUG" = Just "Leu"
  | x == "CUU" = Just "Leu"
  | x == "CUC" = Just "Leu"
  | x == "CUA" = Just "Leu"
  | x == "CUG" = Just "Leu"
  | x == "AUU" = Just "Ile"
  | x == "AUC" = Just "Ile"
  | x == "AUA" = Just "Ile"
  | x == "AUG" = Just "Met"
  | x == "GUU" = Just "Val"
  | x == "GUC" = Just "Val"
  | x == "GUA" = Just "Val"
  | x == "GUG" = Just "Val"
  | x == "GCG" = Just "Ala"
  | x == "GCA" = Just "Ala"
  | x == "GCC" = Just "Ala"
  | x == "GCU" = Just "Ala"
  | x == "ACG" = Just "Thr"
  | x == "ACA" = Just "Thr"
  | x == "ACC" = Just "Thr"
  | x == "ACU" = Just "Thr"
  | x == "CCG" = Just "Pro"
  | x == "CCA" = Just "Pro"
  | x == "CCC" = Just "Pro"
  | x == "CCU" = Just "Pro"
  | x == "UCG" = Just "Ser"
  | x == "UCA" = Just "Ser"
  | x == "UCC" = Just "Ser"
  | x == "UCU" = Just "Ser"
  | x == "UAU" = Just "Tyr"
  | x == "UAC" = Just "Tyr"
  | x == "UAA" = Nothing
  | x == "UAG" = Nothing
  | x == "UGA" = Nothing
  | x == "CAU" = Just "His"
  | x == "CAC" = Just "His"
  | x == "CAA" = Just "Gln"
  | x == "CAG" = Just "Gln"
  | x == "AAU" = Just "Asn"
  | x == "AAC" = Just "Asn"
  | x == "AAA" = Just "Lys"
  | x == "AAG" = Just "Lys"
  | x == "GAU" = Just "Asp"
  | x == "GAC" = Just "Asp"
  | x == "GAA" = Just "Glu"
  | x == "GAG" = Just "Glu"
  | x == "UGU" = Just "Cys"
  | x == "UGC" = Just "Cys"
  | x == "UGG" = Just "Trp"
  | x == "CGU" = Just "Arg"
  | x == "CGC" = Just "Arg"
  | x == "CGA" = Just "Arg"
  | x == "CGG" = Just "Arg"
  | x == "AGU" = Just "Ser"
  | x == "AGC" = Just "Ser"
  | x == "AGA" = Just "Arg"
  | x == "AGG" = Just "Arg"
  | x == "GGU" = Just "Gly"
  | x == "GGC" = Just "Gly"
  | x == "GGA" = Just "Gly"
  | x == "GGG" = Just "Gly"
  | otherwise = Nothing


doStuff :: RNA -> [String]
doStuff (x:y:z:xs) 
  | (makeAminoAcid ([x]++[y]++[z])) == Nothing = []
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Gly" = ["Gly"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Arg" = ["Arg"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Ser" = ["Ser"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Trp" = ["Trp"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Cys" = ["Cys"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Glu" = ["Glu"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Asp" = ["Asp"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Lys" = ["Lys"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Asn" = ["Asn"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Gln" = ["Gln"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "His" = ["His"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Tyr" = ["Tyr"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Pro" = ["Pro"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Thr" = ["Thr"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Ala" = ["Ala"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Val" = ["Val"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Met" = ["Met"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Ile" = ["Ile"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Leu" = ["Leu"] ++ doStuff xs
  | (makeAminoAcid ([x]++[y]++[z])) == Just "Phe" = ["Phe"] ++ doStuff xs
  | otherwise = []
doStuff (y:z:xs) = []
doStuff (x:xs) = []
doStuff [] = []
translateRNA :: RNA -> [String] 
translateRNA (x:y:z:xs) 
  | (makeAminoAcid ([x]++[y]++[z])) == (Just "Met") = doStuff (x:y:z:xs)
  | otherwise = translateRNA (y:z:xs)
