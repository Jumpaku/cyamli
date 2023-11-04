/** Schema - CLI Schema */
type Schema = Program;

/** 
 * Program - Root command that may have a name and a version.
 * It consists of Commands recursively.
 */
type Program = Command & {
    /** 
     * name - the name of your program. 
     * The default value is an empty string.
     */
    name?: string;
    /** 
     * version - the version of your program. 
     * The default value is an empty string. 
     */
    version?: string;
};

/** Command - The root command or subcommand. */
type Command = {
    /** 
     * description - the description of this command. 
     * The default value is an empty string.
     */
    description?: string;
    /** 
     * options - a mapping from option names to Options. 
     * The default value is an empty object.
     * option_name is a name of an option, which must match the regular expression `^(-[a-z][a-z0-9]*)+$` and be unique in options of this Command.
     */
    options?: { [option_name: string]: Option };
    /** 
     * arguments - a list of Arguments. 
     * The default value is an empty array.
     */
    arguments?: Argument[];
    /** 
     * subcommands - a mapping from subcommand names to child Commands. 
     * The default value is an empty object.
     * subcommand_name is a name of a subcommand, which must match the regular expression `^[a-z][a-z0-9]*$` and be unique in subcommands of this Command.
     */
    subcommands?: { [subcommand_name: string]: Command };
}

/** Type - type of values of options or arguments. */
type Type = "boolean" | "string" | "integer" | "float";

/** Option - an optional argument for a Command. */
type Option = {
    /**
     * short - short name of this Option, which must match the regular expression `^-[a-z]$` and be unique in the command which this option belongs to. 
     * If short is not specified then short name for this option is not available.
     */
    short?: string;
    /** 
     * description - the description of this option. 
     * The default value is an empty string. 
     */
    description?: string;
    /** 
     * type - type of the value that is assignable to this option. 
     * The default value is "string".
     */
    type?: Type;
    /**
     * default - string representing the default value of this option.
     * The default value for each types is as follows.
     * - boolean: "false"
     * - string: ""
     * - integer: "0"
     * - float: "0.0"
     */
    default?: string;
};

/** Argument - a positional and required argument for a Command. */
type Argument = {
    /** 
     * name - the name of this argument, which must match the regular expression `^[a-z][a-z0-9]*(_[a-z0-9])*$` and be unique in the command which this argument belongs to. 
     */
    name: string;
    /** 
     * description - the description of this argument. 
     * The default value is an empty string.
     * */
    description?: string;
    /** 
     * type - type of the value that is assignable to this argument. 
     * The default value is "string".
     */
    type?: Type;
    /**
     * variadic - whether this argument is variadic (i.e. can have zero or more values).
     * It can be true only if this argument is the last argument in the belonging Command.
     */
    variadic?: boolean
};