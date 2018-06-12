package org.example.biojava;
// Simple illustrative BioJava-style PDB reader (toy)
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
public class PDBReader {
    public static void main(String[] args) throws IOException {
        if(args.length < 1) { System.err.println("Usage: PDBReader file.pdb"); System.exit(1); }
        try(BufferedReader br = new BufferedReader(new FileReader(args[0]))) {
            String line;
            while((line = br.readLine()) != null) {
                if(line.startsWith("ATOM") || line.startsWith("HETATM")) {
                    System.out.println(line.substring(0,54));
                }
            }
        }
    }
}
