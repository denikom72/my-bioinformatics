package org.example.biojava;
public class SequenceUtils {
    public static double gcContent(String seq) {
        int gc = 0; for(char c: seq.toUpperCase().toCharArray()) if(c=='G' || c=='C') gc++;
        return seq.length()>0 ? (gc*100.0)/seq.length() : 0.0;
    }
}
