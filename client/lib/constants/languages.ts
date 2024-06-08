export const LANGUAGE_VERSIONS = {
    python: "3.10.0",
    java: "15.0.1",
    cpp: "11.2.0",
    javascript: "18.15.8",
}

export const CODE_SNIPPETS = {
    python: `def greet(name):\n\tprint("Hello, " + name + "!")\n\ngreet("Alex")\n`,
    java: `public class Main {\n\tpublic static void main(String[] args) {\n\t\tgreet("Alex");\n\t}\n\n\tpublic static void greet(String name) {\n\t\tSystem.out.println("Hello, " + name + "!");\n\t}\n}\n`,
    cpp: `#include <iostream>\n#include <string>\n\nvoid greet(const std::string& name) {\n\tstd::cout << "Hello, " << name << "!" << std::endl;\n}\n\nint main() {\n\tgreet("Alex");\n\treturn 0;\n}\n`,
    javascript: `function greet(name) {\n\tconsole.log("Hello, " + name + "!");\n}\n\ngreet("Alex");\n`,
}