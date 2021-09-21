 Description:

   pdlw prints the longest substring in a dictionary, containing only specified symbols
   Expected dictionary is map[int]string, expected file format is JSON
   Expected type of symbols file is one line plaintext file

 Usage:
 
    Read and parse dictionary:

       pdlw parse [dict1.json] [--symbols "abcd"]
       pdlw parse [dict1.json] [--symbols-from-file symbols.txt]

   To be implemented:

       pdlw generate dict1.json [--entries 5] [--entries-size-range 15 25]

  --entries N  Specifies number of dictionary entries (min:1, max:100, default:5)

  --entries-size-range X Y
               Specifies minimum and maximum length of dictionary value in symbols 
               ( X<Y, min:1, max:50, default: X=15, Y=25 )

