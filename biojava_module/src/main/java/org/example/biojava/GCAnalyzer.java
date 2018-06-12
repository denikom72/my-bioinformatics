// Developed by J. Müller (DE) 2018
package org.example.biojava;
public class GCAnalyzer {
    public static void main(String[] args) {
        String seq = args.length>0 ? args[0] : "ATGCGCGATAT";
        System.out.println("GC%: " + SequenceUtils.gcContent(seq));
    }
}
