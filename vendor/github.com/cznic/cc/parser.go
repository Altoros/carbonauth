// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2016 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on [0], 6.10. Substantial portions of expression AST size
// optimizations are from [2], license of which follows.

// ----------------------------------------------------------------------------

// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This grammar is derived from the C grammar in the 'ansitize'
// program, which carried this notice:
//
// Copyright (c) 2006 Russ Cox,
// 	Massachusetts Institute of Technology
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the
// Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute,
// sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall
// be included in all copies or substantial portions of the
// Software.
//
// The software is provided "as is", without warranty of any
// kind, express or implied, including but not limited to the
// warranties of merchantability, fitness for a particular
// purpose and noninfringement.  In no event shall the authors
// or copyright holders be liable for any claim, damages or
// other liability, whether in an action of contract, tort or
// otherwise, arising from, out of or in connection with the
// software or the use or other dealings in the software.

package cc

import __yyfmt__ "fmt"

import (
	"fmt"

	"github.com/cznic/golex/lex"
	"github.com/cznic/xc"
)

type yySymType struct {
	yys       int
	Token     xc.Token
	groupPart Node
	node      Node
	toks      PPTokenList
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault           = 57460
	yyEofCode           = 57344
	ADDASSIGN           = 57346
	ALIGNOF             = 57347
	ANDAND              = 57348
	ANDASSIGN           = 57349
	ARROW               = 57350
	ASM                 = 57351
	AUTO                = 57352
	BOOL                = 57353
	BREAK               = 57354
	CASE                = 57355
	CAST                = 57356
	CHAR                = 57357
	CHARCONST           = 57358
	COMPLEX             = 57359
	CONST               = 57360
	CONSTANT_EXPRESSION = 1048577
	CONTINUE            = 57361
	DDD                 = 57362
	DEC                 = 57363
	DEFAULT             = 57364
	DIVASSIGN           = 57365
	DO                  = 57366
	DOUBLE              = 57367
	ELSE                = 57368
	ENUM                = 57369
	EQ                  = 57370
	EXTERN              = 57371
	FLOAT               = 57372
	FLOATCONST          = 57373
	FOR                 = 57374
	GEQ                 = 57375
	GOTO                = 57376
	IDENTIFIER          = 57377
	IDENTIFIER_LPAREN   = 57378
	IDENTIFIER_NONREPL  = 57379
	IF                  = 57380
	INC                 = 57381
	INLINE              = 57382
	INT                 = 57383
	INTCONST            = 57384
	LEQ                 = 57385
	LONG                = 57386
	LONGCHARCONST       = 57387
	LONGSTRINGLITERAL   = 57388
	LSH                 = 57389
	LSHASSIGN           = 57390
	MODASSIGN           = 57391
	MULASSIGN           = 57392
	NEQ                 = 57393
	NOELSE              = 57394
	NORETURN            = 57395
	NOSEMI              = 57396
	ORASSIGN            = 57397
	OROR                = 57398
	PPDEFINE            = 57399
	PPELIF              = 57400
	PPELSE              = 57401
	PPENDIF             = 57402
	PPERROR             = 57403
	PPHASH_NL           = 57404
	PPHEADER_NAME       = 57405
	PPIF                = 57406
	PPIFDEF             = 57407
	PPIFNDEF            = 57408
	PPINCLUDE           = 57409
	PPINCLUDE_NEXT      = 57410
	PPLINE              = 57411
	PPNONDIRECTIVE      = 57412
	PPNUMBER            = 57413
	PPOTHER             = 57414
	PPPASTE             = 57415
	PPPRAGMA            = 57416
	PPUNDEF             = 57417
	PREPROCESSING_FILE  = 1048576
	REGISTER            = 57418
	RESTRICT            = 57419
	RETURN              = 57420
	RSH                 = 57421
	RSHASSIGN           = 57422
	SHORT               = 57423
	SIGNED              = 57424
	SIZEOF              = 57425
	STATIC              = 57426
	STATIC_ASSERT       = 57427
	STRINGLITERAL       = 57428
	STRUCT              = 57429
	SUBASSIGN           = 57430
	SWITCH              = 57431
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57432
	TYPEDEFNAME         = 57433
	TYPEOF              = 57434
	UNARY               = 57435
	UNION               = 57436
	UNSIGNED            = 57437
	VOID                = 57438
	VOLATILE            = 57439
	WHILE               = 57440
	XORASSIGN           = 57441
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -328
)

var (
	yyXLAT = map[int]int{
		40:      0,   // '(' (333x)
		42:      1,   // '*' (294x)
		57377:   2,   // IDENTIFIER (241x)
		38:      3,   // '&' (226x)
		43:      4,   // '+' (226x)
		45:      5,   // '-' (226x)
		57348:   6,   // ANDAND (226x)
		57363:   7,   // DEC (226x)
		57381:   8,   // INC (226x)
		59:      9,   // ';' (220x)
		41:      10,  // ')' (203x)
		44:      11,  // ',' (191x)
		57428:   12,  // STRINGLITERAL (169x)
		91:      13,  // '[' (167x)
		33:      14,  // '!' (150x)
		126:     15,  // '~' (150x)
		57347:   16,  // ALIGNOF (150x)
		57358:   17,  // CHARCONST (150x)
		57373:   18,  // FLOATCONST (150x)
		57384:   19,  // INTCONST (150x)
		57387:   20,  // LONGCHARCONST (150x)
		57388:   21,  // LONGSTRINGLITERAL (150x)
		57425:   22,  // SIZEOF (150x)
		57439:   23,  // VOLATILE (141x)
		57360:   24,  // CONST (139x)
		57419:   25,  // RESTRICT (139x)
		125:     26,  // '}' (132x)
		57353:   27,  // BOOL (129x)
		57357:   28,  // CHAR (129x)
		57359:   29,  // COMPLEX (129x)
		57367:   30,  // DOUBLE (129x)
		57369:   31,  // ENUM (129x)
		57372:   32,  // FLOAT (129x)
		57383:   33,  // INT (129x)
		57386:   34,  // LONG (129x)
		57423:   35,  // SHORT (129x)
		57424:   36,  // SIGNED (129x)
		57429:   37,  // STRUCT (129x)
		57433:   38,  // TYPEDEFNAME (129x)
		57434:   39,  // TYPEOF (129x)
		57436:   40,  // UNION (129x)
		57437:   41,  // UNSIGNED (129x)
		57438:   42,  // VOID (129x)
		58:      43,  // ':' (126x)
		57426:   44,  // STATIC (120x)
		57352:   45,  // AUTO (114x)
		57371:   46,  // EXTERN (114x)
		57382:   47,  // INLINE (114x)
		57395:   48,  // NORETURN (114x)
		57418:   49,  // REGISTER (114x)
		57432:   50,  // TYPEDEF (114x)
		57344:   51,  // $end (106x)
		61:      52,  // '=' (96x)
		123:     53,  // '{' (93x)
		57503:   54,  // Expression (88x)
		93:      55,  // ']' (86x)
		46:      56,  // '.' (85x)
		57351:   57,  // ASM (84x)
		57427:   58,  // STATIC_ASSERT (79x)
		37:      59,  // '%' (77x)
		47:      60,  // '/' (77x)
		60:      61,  // '<' (77x)
		62:      62,  // '>' (77x)
		63:      63,  // '?' (77x)
		94:      64,  // '^' (77x)
		124:     65,  // '|' (77x)
		57346:   66,  // ADDASSIGN (77x)
		57349:   67,  // ANDASSIGN (77x)
		57350:   68,  // ARROW (77x)
		57365:   69,  // DIVASSIGN (77x)
		57370:   70,  // EQ (77x)
		57375:   71,  // GEQ (77x)
		57385:   72,  // LEQ (77x)
		57389:   73,  // LSH (77x)
		57390:   74,  // LSHASSIGN (77x)
		57391:   75,  // MODASSIGN (77x)
		57392:   76,  // MULASSIGN (77x)
		57393:   77,  // NEQ (77x)
		57397:   78,  // ORASSIGN (77x)
		57398:   79,  // OROR (77x)
		57421:   80,  // RSH (77x)
		57422:   81,  // RSHASSIGN (77x)
		57430:   82,  // SUBASSIGN (77x)
		57441:   83,  // XORASSIGN (77x)
		10:      84,  // '\n' (58x)
		57376:   85,  // GOTO (55x)
		57440:   86,  // WHILE (53x)
		57354:   87,  // BREAK (52x)
		57355:   88,  // CASE (52x)
		57361:   89,  // CONTINUE (52x)
		57364:   90,  // DEFAULT (52x)
		57366:   91,  // DO (52x)
		57374:   92,  // FOR (52x)
		57380:   93,  // IF (52x)
		57414:   94,  // PPOTHER (52x)
		57420:   95,  // RETURN (52x)
		57431:   96,  // SWITCH (52x)
		57402:   97,  // PPENDIF (44x)
		57401:   98,  // PPELSE (40x)
		57400:   99,  // PPELIF (39x)
		57399:   100, // PPDEFINE (35x)
		57403:   101, // PPERROR (35x)
		57404:   102, // PPHASH_NL (35x)
		57406:   103, // PPIF (35x)
		57407:   104, // PPIFDEF (35x)
		57408:   105, // PPIFNDEF (35x)
		57409:   106, // PPINCLUDE (35x)
		57410:   107, // PPINCLUDE_NEXT (35x)
		57411:   108, // PPLINE (35x)
		57412:   109, // PPNONDIRECTIVE (35x)
		57416:   110, // PPPRAGMA (35x)
		57417:   111, // PPUNDEF (35x)
		57368:   112, // ELSE (32x)
		57555:   113, // TypeQualifier (28x)
		57504:   114, // ExpressionList (26x)
		57528:   115, // PPTokenList (22x)
		57530:   116, // PPTokens (22x)
		57499:   117, // EnumSpecifier (20x)
		57550:   118, // StructOrUnion (20x)
		57551:   119, // StructOrUnionSpecifier (20x)
		57558:   120, // TypeSpecifier (20x)
		57505:   121, // ExpressionListOpt (18x)
		57470:   122, // BasicAssemblerStatement (15x)
		57476:   123, // CompoundStatement (15x)
		57482:   124, // DeclarationSpecifiers (15x)
		57511:   125, // FunctionSpecifier (15x)
		57529:   126, // PPTokenListOpt (15x)
		57545:   127, // StorageClassSpecifier (15x)
		57468:   128, // AssemblerStatement (13x)
		57507:   129, // ExpressionStatement (12x)
		57525:   130, // IterationStatement (12x)
		57526:   131, // JumpStatement (12x)
		57527:   132, // LabeledStatement (12x)
		57535:   133, // Pointer (12x)
		57539:   134, // SelectionStatement (12x)
		57543:   135, // Statement (12x)
		57536:   136, // PointerOpt (11x)
		57484:   137, // Declarator (9x)
		57544:   138, // StaticAssertDeclaration (9x)
		57478:   139, // ControlLine (8x)
		57514:   140, // GroupPart (8x)
		57518:   141, // IfGroup (8x)
		57519:   142, // IfSection (8x)
		57552:   143, // TextLine (8x)
		57479:   144, // Declaration (7x)
		57454:   145, // $@4 (6x)
		57477:   146, // ConstantExpression (6x)
		57362:   147, // DDD (6x)
		57512:   148, // GroupList (6x)
		57466:   149, // AssemblerOperand (5x)
		57469:   150, // AssemblerSymbolicNameOpt (5x)
		57513:   151, // GroupListOpt (5x)
		57538:   152, // ReplacementList (5x)
		57540:   153, // SpecifierQualifierList (5x)
		57556:   154, // TypeQualifierList (5x)
		57459:   155, // $@9 (4x)
		57461:   156, // AbstractDeclarator (4x)
		57467:   157, // AssemblerOperands (4x)
		57483:   158, // DeclarationSpecifiersOpt (4x)
		57488:   159, // Designator (4x)
		57523:   160, // Initializer (4x)
		57531:   161, // ParameterDeclaration (4x)
		57554:   162, // TypeName (4x)
		57557:   163, // TypeQualifierListOpt (4x)
		57465:   164, // AssemblerInstructions (3x)
		57475:   165, // CommaOpt (3x)
		57486:   166, // Designation (3x)
		57487:   167, // DesignationOpt (3x)
		57489:   168, // DesignatorList (3x)
		57506:   169, // ExpressionOpt (3x)
		57515:   170, // IdentifierList (3x)
		57520:   171, // InitDeclarator (3x)
		57532:   172, // ParameterList (3x)
		57533:   173, // ParameterTypeList (3x)
		57443:   174, // $@10 (2x)
		57447:   175, // $@14 (2x)
		57449:   176, // $@16 (2x)
		57450:   177, // $@17 (2x)
		57451:   178, // $@18 (2x)
		57455:   179, // $@5 (2x)
		57462:   180, // AbstractDeclaratorOpt (2x)
		57471:   181, // BlockItem (2x)
		57474:   182, // Clobbers (2x)
		57481:   183, // DeclarationListOpt (2x)
		57485:   184, // DeclaratorOpt (2x)
		57490:   185, // DirectAbstractDeclarator (2x)
		57491:   186, // DirectAbstractDeclaratorOpt (2x)
		57492:   187, // DirectDeclarator (2x)
		57493:   188, // ElifGroup (2x)
		57500:   189, // EnumerationConstant (2x)
		57501:   190, // Enumerator (2x)
		57508:   191, // ExternalDeclaration (2x)
		57509:   192, // FunctionBody (2x)
		57510:   193, // FunctionDefinition (2x)
		57516:   194, // IdentifierListOpt (2x)
		57517:   195, // IdentifierOpt (2x)
		57521:   196, // InitDeclaratorList (2x)
		57522:   197, // InitDeclaratorListOpt (2x)
		57524:   198, // InitializerList (2x)
		57534:   199, // ParameterTypeListOpt (2x)
		57541:   200, // SpecifierQualifierListOpt (2x)
		57546:   201, // StructDeclaration (2x)
		57548:   202, // StructDeclarator (2x)
		57559:   203, // VolatileOpt (2x)
		57442:   204, // $@1 (1x)
		57444:   205, // $@11 (1x)
		57445:   206, // $@12 (1x)
		57446:   207, // $@13 (1x)
		57448:   208, // $@15 (1x)
		57452:   209, // $@2 (1x)
		57453:   210, // $@3 (1x)
		57456:   211, // $@6 (1x)
		57457:   212, // $@7 (1x)
		57458:   213, // $@8 (1x)
		57463:   214, // ArgumentExpressionList (1x)
		57464:   215, // ArgumentExpressionListOpt (1x)
		57472:   216, // BlockItemList (1x)
		57473:   217, // BlockItemListOpt (1x)
		1048577: 218, // CONSTANT_EXPRESSION (1x)
		57480:   219, // DeclarationList (1x)
		57494:   220, // ElifGroupList (1x)
		57495:   221, // ElifGroupListOpt (1x)
		57496:   222, // ElseGroup (1x)
		57497:   223, // ElseGroupOpt (1x)
		57498:   224, // EndifLine (1x)
		57502:   225, // EnumeratorList (1x)
		57378:   226, // IDENTIFIER_LPAREN (1x)
		1048576: 227, // PREPROCESSING_FILE (1x)
		57537:   228, // PreprocessingFile (1x)
		57542:   229, // Start (1x)
		57547:   230, // StructDeclarationList (1x)
		57549:   231, // StructDeclaratorList (1x)
		1048578: 232, // TRANSLATION_UNIT (1x)
		57553:   233, // TranslationUnit (1x)
		57460:   234, // $default (0x)
		57356:   235, // CAST (0x)
		57345:   236, // error (0x)
		57379:   237, // IDENTIFIER_NONREPL (0x)
		57394:   238, // NOELSE (0x)
		57396:   239, // NOSEMI (0x)
		57405:   240, // PPHEADER_NAME (0x)
		57413:   241, // PPNUMBER (0x)
		57415:   242, // PPPASTE (0x)
		57435:   243, // UNARY (0x)
	}

	yySymNames = []string{
		"'('",
		"'*'",
		"IDENTIFIER",
		"'&'",
		"'+'",
		"'-'",
		"ANDAND",
		"DEC",
		"INC",
		"';'",
		"')'",
		"','",
		"STRINGLITERAL",
		"'['",
		"'!'",
		"'~'",
		"ALIGNOF",
		"CHARCONST",
		"FLOATCONST",
		"INTCONST",
		"LONGCHARCONST",
		"LONGSTRINGLITERAL",
		"SIZEOF",
		"VOLATILE",
		"CONST",
		"RESTRICT",
		"'}'",
		"BOOL",
		"CHAR",
		"COMPLEX",
		"DOUBLE",
		"ENUM",
		"FLOAT",
		"INT",
		"LONG",
		"SHORT",
		"SIGNED",
		"STRUCT",
		"TYPEDEFNAME",
		"TYPEOF",
		"UNION",
		"UNSIGNED",
		"VOID",
		"':'",
		"STATIC",
		"AUTO",
		"EXTERN",
		"INLINE",
		"NORETURN",
		"REGISTER",
		"TYPEDEF",
		"$end",
		"'='",
		"'{'",
		"Expression",
		"']'",
		"'.'",
		"ASM",
		"STATIC_ASSERT",
		"'%'",
		"'/'",
		"'<'",
		"'>'",
		"'?'",
		"'^'",
		"'|'",
		"ADDASSIGN",
		"ANDASSIGN",
		"ARROW",
		"DIVASSIGN",
		"EQ",
		"GEQ",
		"LEQ",
		"LSH",
		"LSHASSIGN",
		"MODASSIGN",
		"MULASSIGN",
		"NEQ",
		"ORASSIGN",
		"OROR",
		"RSH",
		"RSHASSIGN",
		"SUBASSIGN",
		"XORASSIGN",
		"'\\n'",
		"GOTO",
		"WHILE",
		"BREAK",
		"CASE",
		"CONTINUE",
		"DEFAULT",
		"DO",
		"FOR",
		"IF",
		"PPOTHER",
		"RETURN",
		"SWITCH",
		"PPENDIF",
		"PPELSE",
		"PPELIF",
		"PPDEFINE",
		"PPERROR",
		"PPHASH_NL",
		"PPIF",
		"PPIFDEF",
		"PPIFNDEF",
		"PPINCLUDE",
		"PPINCLUDE_NEXT",
		"PPLINE",
		"PPNONDIRECTIVE",
		"PPPRAGMA",
		"PPUNDEF",
		"ELSE",
		"TypeQualifier",
		"ExpressionList",
		"PPTokenList",
		"PPTokens",
		"EnumSpecifier",
		"StructOrUnion",
		"StructOrUnionSpecifier",
		"TypeSpecifier",
		"ExpressionListOpt",
		"BasicAssemblerStatement",
		"CompoundStatement",
		"DeclarationSpecifiers",
		"FunctionSpecifier",
		"PPTokenListOpt",
		"StorageClassSpecifier",
		"AssemblerStatement",
		"ExpressionStatement",
		"IterationStatement",
		"JumpStatement",
		"LabeledStatement",
		"Pointer",
		"SelectionStatement",
		"Statement",
		"PointerOpt",
		"Declarator",
		"StaticAssertDeclaration",
		"ControlLine",
		"GroupPart",
		"IfGroup",
		"IfSection",
		"TextLine",
		"Declaration",
		"$@4",
		"ConstantExpression",
		"DDD",
		"GroupList",
		"AssemblerOperand",
		"AssemblerSymbolicNameOpt",
		"GroupListOpt",
		"ReplacementList",
		"SpecifierQualifierList",
		"TypeQualifierList",
		"$@9",
		"AbstractDeclarator",
		"AssemblerOperands",
		"DeclarationSpecifiersOpt",
		"Designator",
		"Initializer",
		"ParameterDeclaration",
		"TypeName",
		"TypeQualifierListOpt",
		"AssemblerInstructions",
		"CommaOpt",
		"Designation",
		"DesignationOpt",
		"DesignatorList",
		"ExpressionOpt",
		"IdentifierList",
		"InitDeclarator",
		"ParameterList",
		"ParameterTypeList",
		"$@10",
		"$@14",
		"$@16",
		"$@17",
		"$@18",
		"$@5",
		"AbstractDeclaratorOpt",
		"BlockItem",
		"Clobbers",
		"DeclarationListOpt",
		"DeclaratorOpt",
		"DirectAbstractDeclarator",
		"DirectAbstractDeclaratorOpt",
		"DirectDeclarator",
		"ElifGroup",
		"EnumerationConstant",
		"Enumerator",
		"ExternalDeclaration",
		"FunctionBody",
		"FunctionDefinition",
		"IdentifierListOpt",
		"IdentifierOpt",
		"InitDeclaratorList",
		"InitDeclaratorListOpt",
		"InitializerList",
		"ParameterTypeListOpt",
		"SpecifierQualifierListOpt",
		"StructDeclaration",
		"StructDeclarator",
		"VolatileOpt",
		"$@1",
		"$@11",
		"$@12",
		"$@13",
		"$@15",
		"$@2",
		"$@3",
		"$@6",
		"$@7",
		"$@8",
		"ArgumentExpressionList",
		"ArgumentExpressionListOpt",
		"BlockItemList",
		"BlockItemListOpt",
		"CONSTANT_EXPRESSION",
		"DeclarationList",
		"ElifGroupList",
		"ElifGroupListOpt",
		"ElseGroup",
		"ElseGroupOpt",
		"EndifLine",
		"EnumeratorList",
		"IDENTIFIER_LPAREN",
		"PREPROCESSING_FILE",
		"PreprocessingFile",
		"Start",
		"StructDeclarationList",
		"StructDeclaratorList",
		"TRANSLATION_UNIT",
		"TranslationUnit",
		"$default",
		"CAST",
		"error",
		"IDENTIFIER_NONREPL",
		"NOELSE",
		"NOSEMI",
		"PPHEADER_NAME",
		"PPNUMBER",
		"PPPASTE",
		"UNARY",
	}

	yyTokenLiteralStrings = map[int]string{
		57377:   "identifier",
		57348:   "&&",
		57363:   "--",
		57381:   "++",
		57428:   "string literal",
		57347:   "_Alignof",
		57358:   "character constant",
		57373:   "floating-point constant",
		57384:   "integer constant",
		57387:   "long character constant",
		57388:   "long string constant",
		57425:   "sizeof",
		57439:   "volatile",
		57360:   "const",
		57419:   "restrict",
		57353:   "_Bool",
		57357:   "char",
		57359:   "_Complex",
		57367:   "double",
		57369:   "enum",
		57372:   "float",
		57383:   "int",
		57386:   "long",
		57423:   "short",
		57424:   "signed",
		57429:   "struct",
		57433:   "typedefname",
		57434:   "typeof",
		57436:   "union",
		57437:   "unsigned",
		57438:   "void",
		57426:   "static",
		57352:   "auto",
		57371:   "extern",
		57382:   "inline",
		57395:   "_Noreturn",
		57418:   "register",
		57432:   "typedef",
		57351:   "asm",
		57427:   "_Static_assert",
		57346:   "+=",
		57349:   "&=",
		57350:   "->",
		57365:   "/=",
		57370:   "==",
		57375:   ">=",
		57385:   "<=",
		57389:   "<<",
		57390:   "<<=",
		57391:   "%=",
		57392:   "*=",
		57393:   "!=",
		57397:   "|=",
		57398:   "||",
		57421:   ">>",
		57422:   ">>=",
		57430:   "-=",
		57441:   "^=",
		57376:   "goto",
		57440:   "while",
		57354:   "break",
		57355:   "case",
		57361:   "continue",
		57364:   "default",
		57366:   "do",
		57374:   "for",
		57380:   "if",
		57414:   "ppother",
		57420:   "return",
		57431:   "switch",
		57402:   "#endif",
		57401:   "#else",
		57400:   "#elif",
		57399:   "#define",
		57403:   "#error",
		57404:   "#",
		57406:   "#if",
		57407:   "#ifdef",
		57408:   "#ifndef",
		57409:   "#include",
		57410:   "#include_next",
		57411:   "#line",
		57412:   "#foo",
		57416:   "#pragma",
		57417:   "#undef",
		57368:   "else",
		57362:   "...",
		1048577: "constant expression prefix",
		57378:   "identifier immediatelly followed by '('",
		1048576: "preprocessing file prefix",
		1048578: "translation unit prefix",
		57379:   "non replaceable identifier",
		57405:   "header name",
		57413:   "preprocessing number",
		57415:   "##",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {204, 0},
		2:   {229, 3},
		3:   {209, 0},
		4:   {229, 3},
		5:   {210, 0},
		6:   {229, 3},
		7:   {189, 1},
		8:   {214, 1},
		9:   {214, 3},
		10:  {215, 0},
		11:  {215, 1},
		12:  {54, 1},
		13:  {54, 1},
		14:  {54, 1},
		15:  {54, 1},
		16:  {54, 1},
		17:  {54, 1},
		18:  {54, 1},
		19:  {54, 3},
		20:  {54, 4},
		21:  {54, 4},
		22:  {54, 3},
		23:  {54, 3},
		24:  {54, 2},
		25:  {54, 2},
		26:  {54, 7},
		27:  {54, 2},
		28:  {54, 2},
		29:  {54, 2},
		30:  {54, 2},
		31:  {54, 2},
		32:  {54, 2},
		33:  {54, 2},
		34:  {54, 2},
		35:  {54, 2},
		36:  {54, 4},
		37:  {54, 4},
		38:  {54, 3},
		39:  {54, 3},
		40:  {54, 3},
		41:  {54, 3},
		42:  {54, 3},
		43:  {54, 3},
		44:  {54, 3},
		45:  {54, 3},
		46:  {54, 3},
		47:  {54, 3},
		48:  {54, 3},
		49:  {54, 3},
		50:  {54, 3},
		51:  {54, 3},
		52:  {54, 3},
		53:  {54, 3},
		54:  {54, 3},
		55:  {54, 3},
		56:  {54, 5},
		57:  {54, 3},
		58:  {54, 3},
		59:  {54, 3},
		60:  {54, 3},
		61:  {54, 3},
		62:  {54, 3},
		63:  {54, 3},
		64:  {54, 3},
		65:  {54, 3},
		66:  {54, 3},
		67:  {54, 3},
		68:  {54, 4},
		69:  {54, 3},
		70:  {54, 2},
		71:  {54, 4},
		72:  {169, 0},
		73:  {169, 1},
		74:  {114, 1},
		75:  {114, 3},
		76:  {121, 0},
		77:  {121, 1},
		78:  {145, 0},
		79:  {146, 2},
		80:  {144, 3},
		81:  {144, 1},
		82:  {124, 2},
		83:  {124, 2},
		84:  {124, 2},
		85:  {124, 2},
		86:  {158, 0},
		87:  {158, 1},
		88:  {196, 1},
		89:  {196, 3},
		90:  {197, 0},
		91:  {197, 1},
		92:  {171, 1},
		93:  {179, 0},
		94:  {171, 4},
		95:  {127, 1},
		96:  {127, 1},
		97:  {127, 1},
		98:  {127, 1},
		99:  {127, 1},
		100: {120, 1},
		101: {120, 1},
		102: {120, 1},
		103: {120, 1},
		104: {120, 1},
		105: {120, 1},
		106: {120, 1},
		107: {120, 1},
		108: {120, 1},
		109: {120, 1},
		110: {120, 1},
		111: {120, 1},
		112: {120, 1},
		113: {120, 1},
		114: {120, 4},
		115: {120, 4},
		116: {211, 0},
		117: {119, 6},
		118: {119, 2},
		119: {119, 4},
		120: {118, 1},
		121: {118, 1},
		122: {230, 1},
		123: {230, 2},
		124: {201, 3},
		125: {201, 2},
		126: {201, 1},
		127: {153, 2},
		128: {153, 2},
		129: {200, 0},
		130: {200, 1},
		131: {231, 1},
		132: {231, 3},
		133: {202, 1},
		134: {202, 3},
		135: {165, 0},
		136: {165, 1},
		137: {212, 0},
		138: {117, 7},
		139: {117, 2},
		140: {225, 1},
		141: {225, 3},
		142: {190, 1},
		143: {190, 3},
		144: {113, 1},
		145: {113, 1},
		146: {113, 1},
		147: {125, 1},
		148: {125, 1},
		149: {137, 2},
		150: {184, 0},
		151: {184, 1},
		152: {187, 1},
		153: {187, 3},
		154: {187, 5},
		155: {187, 6},
		156: {187, 6},
		157: {187, 5},
		158: {213, 0},
		159: {187, 5},
		160: {187, 4},
		161: {133, 2},
		162: {133, 3},
		163: {136, 0},
		164: {136, 1},
		165: {154, 1},
		166: {154, 2},
		167: {163, 0},
		168: {163, 1},
		169: {173, 1},
		170: {173, 3},
		171: {199, 0},
		172: {199, 1},
		173: {172, 1},
		174: {172, 3},
		175: {161, 2},
		176: {161, 2},
		177: {170, 1},
		178: {170, 3},
		179: {194, 0},
		180: {194, 1},
		181: {195, 0},
		182: {195, 1},
		183: {155, 0},
		184: {162, 3},
		185: {156, 1},
		186: {156, 2},
		187: {180, 0},
		188: {180, 1},
		189: {185, 3},
		190: {185, 4},
		191: {185, 5},
		192: {185, 6},
		193: {185, 6},
		194: {185, 4},
		195: {174, 0},
		196: {185, 4},
		197: {205, 0},
		198: {185, 5},
		199: {186, 0},
		200: {186, 1},
		201: {160, 1},
		202: {160, 4},
		203: {160, 3},
		204: {198, 2},
		205: {198, 4},
		206: {198, 0},
		207: {166, 2},
		208: {167, 0},
		209: {167, 1},
		210: {168, 1},
		211: {168, 2},
		212: {159, 3},
		213: {159, 2},
		214: {135, 1},
		215: {135, 1},
		216: {135, 1},
		217: {135, 1},
		218: {135, 1},
		219: {135, 1},
		220: {135, 1},
		221: {132, 3},
		222: {132, 4},
		223: {132, 3},
		224: {206, 0},
		225: {123, 4},
		226: {216, 1},
		227: {216, 2},
		228: {217, 0},
		229: {217, 1},
		230: {181, 1},
		231: {181, 1},
		232: {129, 2},
		233: {134, 5},
		234: {134, 7},
		235: {134, 5},
		236: {130, 5},
		237: {130, 7},
		238: {130, 9},
		239: {130, 8},
		240: {131, 3},
		241: {131, 2},
		242: {131, 2},
		243: {131, 3},
		244: {131, 3},
		245: {233, 1},
		246: {233, 2},
		247: {191, 1},
		248: {191, 1},
		249: {191, 2},
		250: {191, 1},
		251: {207, 0},
		252: {193, 5},
		253: {175, 0},
		254: {208, 0},
		255: {193, 5},
		256: {176, 0},
		257: {192, 2},
		258: {177, 0},
		259: {192, 3},
		260: {219, 1},
		261: {219, 2},
		262: {183, 0},
		263: {178, 0},
		264: {183, 2},
		265: {164, 1},
		266: {164, 2},
		267: {122, 5},
		268: {203, 0},
		269: {203, 1},
		270: {149, 5},
		271: {157, 1},
		272: {157, 3},
		273: {150, 0},
		274: {150, 3},
		275: {182, 1},
		276: {182, 3},
		277: {128, 1},
		278: {128, 7},
		279: {128, 9},
		280: {128, 11},
		281: {128, 13},
		282: {128, 6},
		283: {128, 8},
		284: {138, 7},
		285: {228, 1},
		286: {148, 1},
		287: {148, 2},
		288: {151, 0},
		289: {151, 1},
		290: {140, 1},
		291: {140, 1},
		292: {140, 3},
		293: {140, 1},
		294: {142, 4},
		295: {141, 4},
		296: {141, 4},
		297: {141, 4},
		298: {220, 1},
		299: {220, 2},
		300: {221, 0},
		301: {221, 1},
		302: {188, 4},
		303: {222, 3},
		304: {223, 0},
		305: {223, 1},
		306: {224, 1},
		307: {139, 3},
		308: {139, 5},
		309: {139, 7},
		310: {139, 5},
		311: {139, 2},
		312: {139, 1},
		313: {139, 3},
		314: {139, 3},
		315: {139, 2},
		316: {139, 3},
		317: {139, 6},
		318: {139, 2},
		319: {139, 4},
		320: {139, 3},
		321: {143, 1},
		322: {152, 1},
		323: {115, 1},
		324: {126, 1},
		325: {126, 2},
		326: {116, 1},
		327: {116, 2},
	}

	yyXErrors = map[yyXError]string{
		{0, 51}:   "invalid empty input",
		{583, -1}: "expected #endif",
		{585, -1}: "expected #endif",
		{1, -1}:   "expected $end",
		{503, -1}: "expected $end",
		{505, -1}: "expected $end",
		{32, -1}:  "expected '('",
		{49, -1}:  "expected '('",
		{75, -1}:  "expected '('",
		{274, -1}: "expected '('",
		{275, -1}: "expected '('",
		{276, -1}: "expected '('",
		{278, -1}: "expected '('",
		{288, -1}: "expected '('",
		{311, -1}: "expected '('",
		{353, -1}: "expected '('",
		{435, -1}: "expected '('",
		{54, -1}:  "expected ')'",
		{79, -1}:  "expected ')'",
		{86, -1}:  "expected ')'",
		{180, -1}: "expected ')'",
		{195, -1}: "expected ')'",
		{198, -1}: "expected ')'",
		{201, -1}: "expected ')'",
		{209, -1}: "expected ')'",
		{214, -1}: "expected ')'",
		{220, -1}: "expected ')'",
		{236, -1}: "expected ')'",
		{241, -1}: "expected ')'",
		{252, -1}: "expected ')'",
		{253, -1}: "expected ')'",
		{343, -1}: "expected ')'",
		{349, -1}: "expected ')'",
		{433, -1}: "expected ')'",
		{489, -1}: "expected ')'",
		{547, -1}: "expected ')'",
		{548, -1}: "expected ')'",
		{555, -1}: "expected ')'",
		{558, -1}: "expected ')'",
		{52, -1}:  "expected ','",
		{267, -1}: "expected ':'",
		{293, -1}: "expected ':'",
		{377, -1}: "expected ':'",
		{479, -1}: "expected ':'",
		{45, -1}:  "expected ';'",
		{55, -1}:  "expected ';'",
		{273, -1}: "expected ';'",
		{280, -1}: "expected ';'",
		{281, -1}: "expected ';'",
		{330, -1}: "expected ';'",
		{339, -1}: "expected ';'",
		{341, -1}: "expected ';'",
		{347, -1}: "expected ';'",
		{356, -1}: "expected ';'",
		{380, -1}: "expected ';'",
		{448, -1}: "expected ';'",
		{387, -1}: "expected '='",
		{91, -1}:  "expected '['",
		{527, -1}: "expected '\\n'",
		{531, -1}: "expected '\\n'",
		{535, -1}: "expected '\\n'",
		{538, -1}: "expected '\\n'",
		{540, -1}: "expected '\\n'",
		{562, -1}: "expected '\\n'",
		{567, -1}: "expected '\\n'",
		{570, -1}: "expected '\\n'",
		{577, -1}: "expected '\\n'",
		{582, -1}: "expected '\\n'",
		{588, -1}: "expected '\\n'",
		{97, -1}:  "expected ']'",
		{188, -1}: "expected ']'",
		{232, -1}: "expected ']'",
		{299, -1}: "expected ']'",
		{401, -1}: "expected ']'",
		{452, -1}: "expected '{'",
		{454, -1}: "expected '{'",
		{466, -1}: "expected '{'",
		{268, -1}: "expected '}'",
		{407, -1}: "expected '}'",
		{423, -1}: "expected '}'",
		{463, -1}: "expected '}'",
		{0, -1}:   "expected Start or one of [constant expression prefix, preprocessing file prefix, translation unit prefix]",
		{208, -1}: "expected abstract declarator or declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{90, -1}:  "expected abstract declarator or optional parameter type list or one of ['(', ')', '*', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{287, -1}: "expected assembler instructions or string literal",
		{289, -1}: "expected assembler instructions or string literal",
		{436, -1}: "expected assembler instructions or string literal",
		{301, -1}: "expected assembler operand or one of ['[', string literal]",
		{317, -1}: "expected assembler operands or one of [')', ':', '[', string literal]",
		{294, -1}: "expected assembler operands or one of ['[', string literal]",
		{320, -1}: "expected assembler operands or one of ['[', string literal]",
		{324, -1}: "expected assembler operands or one of ['[', string literal]",
		{447, -1}: "expected assembler statement or asm",
		{270, -1}: "expected block item or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{302, -1}: "expected clobbers or string literal",
		{327, -1}: "expected clobbers or string literal",
		{446, -1}: "expected compound statement or '{'",
		{64, -1}:  "expected compound statement or expression list or type name or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{249, -1}: "expected compound statement or expression list or type name or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{3, -1}:   "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{50, -1}:  "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{266, -1}: "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{398, -1}: "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{460, -1}: "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{480, -1}: "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{502, -1}: "expected constant expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{440, -1}: "expected declaration list or one of [_Bool, _Complex, _Noreturn, _Static_assert, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{442, -1}: "expected declaration or one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{338, -1}: "expected declaration or optional expression list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{47, -1}:  "expected declarator or one of ['(', '*', identifier]",
		{386, -1}: "expected declarator or one of ['(', '*', identifier]",
		{200, -1}: "expected declarator or optional abstract declarator or one of ['(', ')', '*', ',', '[', identifier]",
		{7, -1}:   "expected declarator or optional init declarator list or one of ['(', '*', ';', identifier]",
		{395, -1}: "expected designator or one of ['.', '=', '[']",
		{203, -1}: "expected direct abstract declarator or direct declarator or one of ['(', '[', identifier]",
		{87, -1}:  "expected direct abstract declarator or one of ['(', '[']",
		{384, -1}: "expected direct declarator or one of ['(', identifier]",
		{575, -1}: "expected elif group or one of [#elif, #else, #endif]",
		{581, -1}: "expected endif line or #endif",
		{512, -1}: "expected endif line or optional elif group list or optional else group or one of [#elif, #else, #endif]",
		{573, -1}: "expected endif line or optional else group or one of [#else, #endif]",
		{455, -1}: "expected enumerator list or identifier",
		{462, -1}: "expected enumerator or one of ['}', identifier]",
		{126, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', ':', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{102, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{354, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{358, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{362, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{366, -1}: "expected expression list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{419, -1}: "expected expression or one of [!=, $end, %=, &&, &=, '!', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '{', '|', '}', '~', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal, |=, ||]",
		{94, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{231, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{434, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{51, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{66, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{67, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{68, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{69, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{70, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{71, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{72, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{73, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{74, -1}:  "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{100, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{108, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{109, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{110, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{111, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{112, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{113, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{114, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{115, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{116, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{117, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{118, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{119, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{120, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{121, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{122, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{123, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{124, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{125, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{127, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{128, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{129, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{130, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{131, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{132, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{133, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{134, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{135, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{136, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{137, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{152, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{154, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{155, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{182, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{189, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{225, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{228, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{279, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{312, -1}: "expected expression or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{98, -1}:  "expected expression or optional type qualifier list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{223, -1}: "expected expression or optional type qualifier list or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{487, -1}: "expected expression or type name or one of [&&, '!', '&', '(', '*', '+', '-', '~', ++, --, _Alignof, _Bool, _Complex, char, character constant, const, double, enum, float, floating-point constant, identifier, int, integer constant, long, long character constant, long string constant, restrict, short, signed, sizeof, string literal, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{6, -1}:   "expected external declaration or one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{439, -1}: "expected function body or one of ['{', asm]",
		{444, -1}: "expected function body or one of ['{', asm]",
		{498, -1}: "expected function body or one of ['{', asm]",
		{499, -1}: "expected function body or one of ['{', asm]",
		{497, -1}: "expected function body or optional declaration list or one of [',', ';', '=', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{438, -1}: "expected function body or optional declaration list or one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{564, -1}: "expected group part or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{506, -1}: "expected group part or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{76, -1}:  "expected identifier",
		{104, -1}: "expected identifier",
		{105, -1}: "expected identifier",
		{217, -1}: "expected identifier",
		{298, -1}: "expected identifier",
		{399, -1}: "expected identifier",
		{514, -1}: "expected identifier",
		{515, -1}: "expected identifier",
		{522, -1}: "expected identifier",
		{306, -1}: "expected identifier list or identifier",
		{544, -1}: "expected identifier list or optional identifier list or one of [')', ..., identifier]",
		{413, -1}: "expected init declarator or one of ['(', '*', identifier]",
		{392, -1}: "expected initializer list or optional comma or one of [&&, '!', '&', '(', '*', '+', ',', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{421, -1}: "expected initializer list or optional comma or one of [&&, '!', '&', '(', '*', '+', ',', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{388, -1}: "expected initializer or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{394, -1}: "expected initializer or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{409, -1}: "expected initializer or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{411, -1}: "expected initializer or one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{406, -1}: "expected initializer or optional designation or one of [&&, '!', '&', '(', '*', '+', '-', '.', '[', '{', '}', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{57, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{58, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{59, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{60, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{61, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{62, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{63, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{77, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{106, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{107, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{139, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{140, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{141, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{142, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{143, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{144, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{145, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{146, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{147, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{148, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{149, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{153, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{157, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{158, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{159, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{160, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{161, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{162, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{163, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{164, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{165, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{166, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{167, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{168, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{169, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{170, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{171, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{172, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{173, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{174, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{175, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{176, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{177, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{181, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{185, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{193, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{248, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{250, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{418, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{420, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{424, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{425, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{426, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{427, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{428, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{429, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{430, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{431, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{432, -1}: "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{65, -1}:  "expected one of [!=, $end, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{150, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{156, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{178, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{183, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{313, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{488, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', ')', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{389, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{256, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{390, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', '}', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{334, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{335, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', ';', '<', '=', '>', '?', '[', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{93, -1}:  "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{101, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{190, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{226, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{229, -1}: "expected one of [!=, %=, &&, &=, '%', '&', '(', '*', '+', '-', '.', '/', '<', '=', '>', '?', '[', ']', '^', '|', *=, ++, +=, --, -=, ->, /=, <<, <<=, <=, ==, >=, >>, >>=, ^=, |=, ||]",
		{507, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{508, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{509, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{511, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{518, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{524, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{526, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{529, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{532, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{534, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{536, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{537, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{539, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{541, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{542, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{545, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{550, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{551, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{553, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{557, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{560, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{561, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{566, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{586, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{587, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{589, -1}: "expected one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, $end, '\\n', ppother]",
		{565, -1}: "expected one of [#elif, #else, #endif]",
		{569, -1}: "expected one of [#elif, #else, #endif]",
		{572, -1}: "expected one of [#elif, #else, #endif]",
		{574, -1}: "expected one of [#elif, #else, #endif]",
		{579, -1}: "expected one of [#elif, #else, #endif]",
		{580, -1}: "expected one of [#elif, #else, #endif]",
		{374, -1}: "expected one of [$end, &&, '!', '&', '(', ')', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{8, -1}:   "expected one of [$end, &&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{56, -1}:  "expected one of [$end, &&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{415, -1}: "expected one of [$end, &&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{42, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{43, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{44, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{46, -1}:  "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{445, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{449, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{450, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{451, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{500, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{501, -1}: "expected one of [$end, '(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{37, -1}:  "expected one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{38, -1}:  "expected one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{39, -1}:  "expected one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', ':', ';', '[', ']', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, auto, char, character constant, const, double, enum, extern, float, floating-point constant, identifier, inline, int, integer constant, long, long character constant, long string constant, register, restrict, short, signed, sizeof, static, string literal, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{95, -1}:  "expected one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{186, -1}: "expected one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', '[', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{259, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{260, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{261, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{262, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{263, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{264, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{265, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{284, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{308, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{316, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{319, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{322, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{323, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{326, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{329, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{331, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{332, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{333, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{336, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{337, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{345, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{351, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{357, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{361, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{365, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{369, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{371, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{372, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{376, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{379, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{417, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, else, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{269, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{271, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{272, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{373, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{396, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{403, -1}: "expected one of [&&, '!', '&', '(', '*', '+', '-', '{', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{453, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{467, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', '{', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{18, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{19, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{20, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{21, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{22, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{23, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{24, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{25, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{26, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{27, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{28, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{29, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{30, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{31, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{464, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{470, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{485, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{490, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{491, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{13, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{14, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{15, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{16, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{17, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{40, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{41, -1}:  "expected one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{492, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{493, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{494, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{495, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{496, -1}: "expected one of ['(', ')', '*', ',', ';', '[', identifier]",
		{245, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		{246, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		{247, -1}: "expected one of ['(', ')', '*', ':', ';', '[', identifier]",
		{206, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{207, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{210, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{219, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{221, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{227, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{230, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{233, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{234, -1}: "expected one of ['(', ')', ',', ':', ';', '=', '[', '{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{85, -1}:  "expected one of ['(', ')', ',', '[', identifier]",
		{244, -1}: "expected one of ['(', ')', ',', '[', identifier]",
		{89, -1}:  "expected one of ['(', ')', ',', '[']",
		{138, -1}: "expected one of ['(', ')', ',', '[']",
		{187, -1}: "expected one of ['(', ')', ',', '[']",
		{191, -1}: "expected one of ['(', ')', ',', '[']",
		{192, -1}: "expected one of ['(', ')', ',', '[']",
		{194, -1}: "expected one of ['(', ')', ',', '[']",
		{202, -1}: "expected one of ['(', ')', ',', '[']",
		{238, -1}: "expected one of ['(', ')', ',', '[']",
		{242, -1}: "expected one of ['(', ')', ',', '[']",
		{285, -1}: "expected one of ['(', goto]",
		{286, -1}: "expected one of ['(', goto]",
		{385, -1}: "expected one of ['(', identifier]",
		{296, -1}: "expected one of [')', ',', ':']",
		{303, -1}: "expected one of [')', ',', ':']",
		{309, -1}: "expected one of [')', ',', ':']",
		{310, -1}: "expected one of [')', ',', ':']",
		{314, -1}: "expected one of [')', ',', ':']",
		{318, -1}: "expected one of [')', ',', ':']",
		{325, -1}: "expected one of [')', ',', ':']",
		{257, -1}: "expected one of [')', ',', ';']",
		{215, -1}: "expected one of [')', ',', ...]",
		{218, -1}: "expected one of [')', ',', ...]",
		{546, -1}: "expected one of [')', ',', ...]",
		{88, -1}:  "expected one of [')', ',']",
		{179, -1}: "expected one of [')', ',']",
		{197, -1}: "expected one of [')', ',']",
		{199, -1}: "expected one of [')', ',']",
		{204, -1}: "expected one of [')', ',']",
		{205, -1}: "expected one of [')', ',']",
		{216, -1}: "expected one of [')', ',']",
		{237, -1}: "expected one of [')', ',']",
		{251, -1}: "expected one of [')', ',']",
		{307, -1}: "expected one of [')', ',']",
		{321, -1}: "expected one of [')', ',']",
		{328, -1}: "expected one of [')', ',']",
		{355, -1}: "expected one of [')', ',']",
		{359, -1}: "expected one of [')', ',']",
		{363, -1}: "expected one of [')', ',']",
		{367, -1}: "expected one of [')', ',']",
		{290, -1}: "expected one of [')', ':', string literal]",
		{292, -1}: "expected one of [')', ':', string literal]",
		{315, -1}: "expected one of [')', ':', string literal]",
		{437, -1}: "expected one of [')', string literal]",
		{478, -1}: "expected one of [',', ':', ';']",
		{151, -1}: "expected one of [',', ':']",
		{297, -1}: "expected one of [',', ':']",
		{304, -1}: "expected one of [',', ':']",
		{383, -1}: "expected one of [',', ';', '=']",
		{408, -1}: "expected one of [',', ';', '}']",
		{412, -1}: "expected one of [',', ';', '}']",
		{381, -1}: "expected one of [',', ';']",
		{382, -1}: "expected one of [',', ';']",
		{391, -1}: "expected one of [',', ';']",
		{414, -1}: "expected one of [',', ';']",
		{475, -1}: "expected one of [',', ';']",
		{477, -1}: "expected one of [',', ';']",
		{481, -1}: "expected one of [',', ';']",
		{484, -1}: "expected one of [',', ';']",
		{456, -1}: "expected one of [',', '=', '}']",
		{459, -1}: "expected one of [',', '=', '}']",
		{184, -1}: "expected one of [',', ']']",
		{405, -1}: "expected one of [',', '}']",
		{410, -1}: "expected one of [',', '}']",
		{458, -1}: "expected one of [',', '}']",
		{461, -1}: "expected one of [',', '}']",
		{465, -1}: "expected one of [',', '}']",
		{397, -1}: "expected one of ['.', '=', '[']",
		{400, -1}: "expected one of ['.', '=', '[']",
		{402, -1}: "expected one of ['.', '=', '[']",
		{404, -1}: "expected one of ['.', '=', '[']",
		{291, -1}: "expected one of [':', string literal]",
		{516, -1}: "expected one of ['\\n', identifier, identifier immediatelly followed by '(']",
		{525, -1}: "expected one of ['\\n', ppother]",
		{528, -1}: "expected one of ['\\n', ppother]",
		{530, -1}: "expected one of ['\\n', ppother]",
		{441, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{443, -1}: "expected one of ['{', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{34, -1}:  "expected one of ['{', identifier]",
		{35, -1}:  "expected one of ['{', identifier]",
		{472, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{474, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{476, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{482, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{486, -1}: "expected one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{554, -1}: "expected one of [..., identifier]",
		{83, -1}:  "expected optional abstract declarator or one of ['(', ')', '*', '[']",
		{103, -1}: "expected optional argument expression list or one of [&&, '!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{254, -1}: "expected optional block item list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{255, -1}: "expected optional block item list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '}', '~', ++, --, _Alignof, _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, break, case, char, character constant, const, continue, default, do, double, enum, extern, float, floating-point constant, for, goto, identifier, if, inline, int, integer constant, long, long character constant, long string constant, register, restrict, return, short, signed, sizeof, static, string literal, struct, switch, typedef, typedefname, typeof, union, unsigned, void, volatile, while]",
		{393, -1}: "expected optional comma or one of [',', '}']",
		{422, -1}: "expected optional comma or one of [',', '}']",
		{457, -1}: "expected optional comma or one of [',', '}']",
		{9, -1}:   "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{10, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{11, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{12, -1}:  "expected optional declaration specifiers or one of ['(', ')', '*', ',', ';', '[', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{342, -1}: "expected optional expression list or one of [&&, '!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{348, -1}: "expected optional expression list or one of [&&, '!', '&', '(', ')', '*', '+', '-', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{282, -1}: "expected optional expression list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{340, -1}: "expected optional expression list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{346, -1}: "expected optional expression list or one of [&&, '!', '&', '(', '*', '+', '-', ';', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{222, -1}: "expected optional expression or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, floating-point constant, identifier, integer constant, long character constant, long string constant, sizeof, string literal]",
		{211, -1}: "expected optional expression or optional type qualifier list or type qualifier list or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{92, -1}:  "expected optional expression or type qualifier list or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{96, -1}:  "expected optional expression or type qualifier or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{563, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{568, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{571, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{578, -1}: "expected optional group list or one of [#, #define, #elif, #else, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{584, -1}: "expected optional group list or one of [#, #define, #endif, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{212, -1}: "expected optional identifier list or parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{33, -1}:  "expected optional identifier or one of ['{', identifier]",
		{36, -1}:  "expected optional identifier or one of ['{', identifier]",
		{258, -1}: "expected optional init declarator list or one of ['(', '*', ';', identifier]",
		{196, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{239, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{240, -1}: "expected optional parameter type list or one of [')', _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{81, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{82, -1}:  "expected optional specifier qualifier list or one of ['(', ')', '*', ':', ';', '[', _Bool, _Complex, char, const, double, enum, float, identifier, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{517, -1}: "expected optional token list or one of ['\\n', ppother]",
		{521, -1}: "expected optional token list or one of ['\\n', ppother]",
		{84, -1}:  "expected optional type qualifier list or pointer or one of ['(', ')', '*', ',', '[', const, identifier, restrict, volatile]",
		{283, -1}: "expected optional volatile or one of ['(', goto, volatile]",
		{48, -1}:  "expected optional volatile or one of ['(', volatile]",
		{235, -1}: "expected parameter declaration or one of [..., _Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{213, -1}: "expected parameter type list or one of [_Bool, _Complex, _Noreturn, auto, char, const, double, enum, extern, float, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{243, -1}: "expected pointer or one of ['(', ')', '*', ',', '[', identifier]",
		{2, -1}:   "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{504, -1}: "expected preprocessing file or one of [#, #define, #error, #foo, #if, #ifdef, #ifndef, #include, #include_next, #line, #pragma, #undef, '\\n', ppother]",
		{543, -1}: "expected replacement list or one of ['\\n', ppother]",
		{549, -1}: "expected replacement list or one of ['\\n', ppother]",
		{552, -1}: "expected replacement list or one of ['\\n', ppother]",
		{556, -1}: "expected replacement list or one of ['\\n', ppother]",
		{559, -1}: "expected replacement list or one of ['\\n', ppother]",
		{80, -1}:  "expected specifier qualifier list or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{277, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{344, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{350, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{360, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{364, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{368, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{370, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{375, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{378, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{416, -1}: "expected statement or one of [&&, '!', '&', '(', '*', '+', '-', ';', '{', '~', ++, --, _Alignof, asm, break, case, character constant, continue, default, do, floating-point constant, for, goto, identifier, if, integer constant, long character constant, long string constant, return, sizeof, string literal, switch, while]",
		{53, -1}:  "expected string literal",
		{295, -1}: "expected string literal",
		{300, -1}: "expected string literal",
		{305, -1}: "expected string literal",
		{468, -1}: "expected struct declaration list or one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{469, -1}: "expected struct declaration list or one of [_Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{471, -1}: "expected struct declaration or one of ['}', _Bool, _Complex, _Static_assert, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{473, -1}: "expected struct declarator list or one of ['(', '*', ':', ';', identifier]",
		{483, -1}: "expected struct declarator or one of ['(', '*', ':', identifier]",
		{533, -1}: "expected token list or one of ['\\n', ppother]",
		{510, -1}: "expected token list or ppother",
		{513, -1}: "expected token list or ppother",
		{519, -1}: "expected token list or ppother",
		{520, -1}: "expected token list or ppother",
		{523, -1}: "expected token list or ppother",
		{576, -1}: "expected token list or ppother",
		{4, -1}:   "expected translation unit or one of ['(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{5, -1}:   "expected translation unit or one of ['(', '*', ';', _Bool, _Complex, _Noreturn, _Static_assert, asm, auto, char, const, double, enum, extern, float, identifier, inline, int, long, register, restrict, short, signed, static, struct, typedef, typedefname, typeof, union, unsigned, void, volatile]",
		{78, -1}:  "expected type name or one of [_Bool, _Complex, char, const, double, enum, float, int, long, restrict, short, signed, struct, typedefname, typeof, union, unsigned, void, volatile]",
		{99, -1}:  "expected type qualifier or one of [&&, '!', '&', '(', ')', '*', '+', ',', '-', '[', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, string literal, volatile]",
		{224, -1}: "expected type qualifier or one of [&&, '!', '&', '(', '*', '+', '-', ']', '~', ++, --, _Alignof, character constant, const, floating-point constant, identifier, integer constant, long character constant, long string constant, restrict, sizeof, static, string literal, volatile]",
		{352, -1}: "expected while",
		{3, 51}:   "unexpected EOF",
		{2, 51}:   "unexpected EOF",
		{4, 51}:   "unexpected EOF",
	}

	yyParseTab = [590][]uint16{
		// 0
		{218: 331, 227: 330, 229: 329, 232: 332},
		{51: 328},
		{84: 327, 94: 327, 100: 327, 327, 327, 327, 327, 327, 327, 327, 327, 327, 327, 327, 204: 832},
		{325, 325, 325, 325, 325, 325, 325, 325, 325, 12: 325, 14: 325, 325, 325, 325, 325, 325, 325, 325, 325, 209: 830},
		{323, 323, 323, 9: 323, 23: 323, 323, 323, 27: 323, 323, 323, 323, 323, 323, 323, 323, 323, 323, 323, 323, 323, 323, 323, 323, 44: 323, 323, 323, 323, 323, 323, 323, 57: 323, 323, 210: 333},
		// 5
		{75, 75, 75, 9: 374, 23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 57: 376, 377, 113: 339, 117: 358, 361, 357, 338, 122: 373, 124: 335, 340, 127: 337, 138: 336, 144: 372, 175: 375, 191: 370, 193: 371, 233: 334},
		{75, 75, 75, 9: 374, 23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 322, 57: 376, 377, 113: 339, 117: 358, 361, 357, 338, 122: 373, 124: 335, 340, 127: 337, 138: 336, 144: 372, 175: 375, 191: 829, 193: 371},
		{165, 412, 165, 9: 238, 133: 713, 136: 712, 825, 171: 709, 196: 710, 708},
		{247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 12: 247, 14: 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 44: 247, 247, 247, 247, 247, 247, 247, 247, 53: 247, 57: 247, 247, 85: 247, 247, 247, 247, 247, 247, 247, 247, 247, 95: 247, 247},
		{242, 242, 242, 9: 242, 242, 242, 13: 242, 23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 113: 339, 117: 358, 361, 357, 338, 124: 821, 340, 127: 337, 158: 824},
		// 10
		{242, 242, 242, 9: 242, 242, 242, 13: 242, 23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 113: 339, 117: 358, 361, 357, 338, 124: 821, 340, 127: 337, 158: 823},
		{242, 242, 242, 9: 242, 242, 242, 13: 242, 23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 113: 339, 117: 358, 361, 357, 338, 124: 821, 340, 127: 337, 158: 822},
		{242, 242, 242, 9: 242, 242, 242, 13: 242, 23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 113: 339, 117: 358, 361, 357, 338, 124: 821, 340, 127: 337, 158: 820},
		{233, 233, 233, 9: 233, 233, 233, 13: 233, 23: 233, 233, 233, 27: 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 44: 233, 233, 233, 233, 233, 233, 233},
		{232, 232, 232, 9: 232, 232, 232, 13: 232, 23: 232, 232, 232, 27: 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 44: 232, 232, 232, 232, 232, 232, 232},
		// 15
		{231, 231, 231, 9: 231, 231, 231, 13: 231, 23: 231, 231, 231, 27: 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 44: 231, 231, 231, 231, 231, 231, 231},
		{230, 230, 230, 9: 230, 230, 230, 13: 230, 23: 230, 230, 230, 27: 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 44: 230, 230, 230, 230, 230, 230, 230},
		{229, 229, 229, 9: 229, 229, 229, 13: 229, 23: 229, 229, 229, 27: 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 44: 229, 229, 229, 229, 229, 229, 229},
		{228, 228, 228, 9: 228, 228, 228, 13: 228, 23: 228, 228, 228, 27: 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228},
		{227, 227, 227, 9: 227, 227, 227, 13: 227, 23: 227, 227, 227, 27: 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227},
		// 20
		{226, 226, 226, 9: 226, 226, 226, 13: 226, 23: 226, 226, 226, 27: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226},
		{225, 225, 225, 9: 225, 225, 225, 13: 225, 23: 225, 225, 225, 27: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225},
		{224, 224, 224, 9: 224, 224, 224, 13: 224, 23: 224, 224, 224, 27: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224},
		{223, 223, 223, 9: 223, 223, 223, 13: 223, 23: 223, 223, 223, 27: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223},
		{222, 222, 222, 9: 222, 222, 222, 13: 222, 23: 222, 222, 222, 27: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222},
		// 25
		{221, 221, 221, 9: 221, 221, 221, 13: 221, 23: 221, 221, 221, 27: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221},
		{220, 220, 220, 9: 220, 220, 220, 13: 220, 23: 220, 220, 220, 27: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220},
		{219, 219, 219, 9: 219, 219, 219, 13: 219, 23: 219, 219, 219, 27: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219},
		{218, 218, 218, 9: 218, 218, 218, 13: 218, 23: 218, 218, 218, 27: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218},
		{217, 217, 217, 9: 217, 217, 217, 13: 217, 23: 217, 217, 217, 27: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217},
		// 30
		{216, 216, 216, 9: 216, 216, 216, 13: 216, 23: 216, 216, 216, 27: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216},
		{215, 215, 215, 9: 215, 215, 215, 13: 215, 23: 215, 215, 215, 27: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215},
		{815},
		{2: 795, 53: 147, 195: 794},
		{2: 208, 53: 208},
		// 35
		{2: 207, 53: 207},
		{2: 781, 53: 147, 195: 780},
		{184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 27: 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 55: 184},
		{183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 27: 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 183, 55: 183},
		{182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 27: 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 182, 55: 182},
		// 40
		{181, 181, 181, 9: 181, 181, 181, 13: 181, 23: 181, 181, 181, 27: 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 181, 44: 181, 181, 181, 181, 181, 181, 181},
		{180, 180, 180, 9: 180, 180, 180, 13: 180, 23: 180, 180, 180, 27: 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 180, 44: 180, 180, 180, 180, 180, 180, 180},
		{83, 83, 83, 9: 83, 23: 83, 83, 83, 27: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 44: 83, 83, 83, 83, 83, 83, 83, 83, 57: 83, 83},
		{81, 81, 81, 9: 81, 23: 81, 81, 81, 27: 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 81, 44: 81, 81, 81, 81, 81, 81, 81, 81, 57: 81, 81},
		{80, 80, 80, 9: 80, 23: 80, 80, 80, 27: 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 44: 80, 80, 80, 80, 80, 80, 80, 80, 57: 80, 80},
		// 45
		{9: 779},
		{78, 78, 78, 9: 78, 23: 78, 78, 78, 27: 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 44: 78, 78, 78, 78, 78, 78, 78, 78, 57: 78, 78},
		{165, 412, 165, 133: 713, 136: 712, 766},
		{60, 23: 614, 203: 763},
		{378},
		// 50
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 379, 380},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 393},
		{11: 381},
		{12: 382},
		{10: 383},
		// 55
		{9: 384},
		{44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 12: 44, 14: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44: 44, 44, 44, 44, 44, 44, 44, 44, 53: 44, 57: 44, 44, 85: 44, 44, 44, 44, 44, 44, 44, 44, 44, 95: 44, 44},
		{316, 316, 3: 316, 316, 316, 316, 316, 316, 316, 316, 316, 13: 316, 26: 316, 43: 316, 51: 316, 316, 55: 316, 316, 59: 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316},
		{315, 315, 3: 315, 315, 315, 315, 315, 315, 315, 315, 315, 13: 315, 26: 315, 43: 315, 51: 315, 315, 55: 315, 315, 59: 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315, 315},
		{314, 314, 3: 314, 314, 314, 314, 314, 314, 314, 314, 314, 13: 314, 26: 314, 43: 314, 51: 314, 314, 55: 314, 314, 59: 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314, 314},
		// 60
		{313, 313, 3: 313, 313, 313, 313, 313, 313, 313, 313, 313, 13: 313, 26: 313, 43: 313, 51: 313, 313, 55: 313, 313, 59: 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 313},
		{312, 312, 3: 312, 312, 312, 312, 312, 312, 312, 312, 312, 13: 312, 26: 312, 43: 312, 51: 312, 312, 55: 312, 312, 59: 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312, 312},
		{311, 311, 3: 311, 311, 311, 311, 311, 311, 311, 311, 311, 13: 311, 26: 311, 43: 311, 51: 311, 311, 55: 311, 311, 59: 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311, 311},
		{310, 310, 3: 310, 310, 310, 310, 310, 310, 310, 310, 310, 13: 310, 26: 310, 43: 310, 51: 310, 310, 55: 310, 310, 59: 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310, 310},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 145, 145, 145, 27: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 53: 582, 478, 114: 579, 123: 581, 155: 408, 162: 761},
		// 65
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 249, 11: 249, 13: 430, 26: 249, 43: 249, 51: 249, 455, 55: 249, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 760},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 759},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 758},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 521},
		// 70
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 757},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 756},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 755},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 754},
		{577, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 578},
		// 75
		{406},
		{2: 405},
		{258, 258, 3: 258, 258, 258, 258, 258, 258, 258, 258, 258, 13: 258, 26: 258, 43: 258, 51: 258, 258, 55: 258, 258, 59: 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258, 258},
		{23: 145, 145, 145, 27: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 155: 408, 162: 407},
		{10: 576},
		// 80
		{23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 113: 410, 117: 358, 361, 357, 409, 153: 411},
		{199, 199, 199, 9: 199, 199, 13: 199, 23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 199, 113: 410, 117: 358, 361, 357, 409, 153: 574, 200: 575},
		{199, 199, 199, 9: 199, 199, 13: 199, 23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 199, 113: 410, 117: 358, 361, 357, 409, 153: 574, 200: 573},
		{165, 412, 10: 141, 13: 165, 133: 413, 136: 415, 156: 416, 180: 414},
		{161, 161, 161, 10: 161, 161, 13: 161, 23: 367, 365, 366, 113: 423, 154: 427, 163: 571},
		// 85
		{164, 2: 164, 10: 143, 143, 13: 164},
		{10: 144},
		{418, 13: 129, 185: 417, 419},
		{10: 140, 140},
		{567, 10: 142, 142, 13: 128},
		// 90
		{165, 412, 10: 133, 13: 165, 23: 133, 133, 133, 27: 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 44: 133, 133, 133, 133, 133, 133, 133, 133: 413, 136: 415, 156: 523, 174: 524},
		{13: 420},
		{392, 422, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 367, 365, 366, 44: 426, 54: 421, 256, 113: 423, 154: 424, 169: 425},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 13: 430, 52: 455, 55: 255, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 521, 522},
		// 95
		{163, 163, 163, 163, 163, 163, 163, 163, 163, 10: 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 44: 163, 55: 163},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 367, 365, 366, 44: 517, 54: 421, 256, 113: 514, 169: 516},
		{55: 515},
		{161, 161, 161, 161, 161, 161, 161, 161, 161, 12: 161, 14: 161, 161, 161, 161, 161, 161, 161, 161, 161, 367, 365, 366, 113: 423, 154: 427, 163: 428},
		{160, 160, 160, 160, 160, 160, 160, 160, 160, 10: 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 160, 367, 365, 366, 113: 514},
		// 100
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 429},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 13: 430, 52: 455, 55: 466, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 478, 114: 512},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 10: 318, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 506, 214: 507, 508},
		{2: 505},
		// 105
		{2: 504},
		{304, 304, 3: 304, 304, 304, 304, 304, 304, 304, 304, 304, 13: 304, 26: 304, 43: 304, 51: 304, 304, 55: 304, 304, 59: 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304, 304},
		{303, 303, 3: 303, 303, 303, 303, 303, 303, 303, 303, 303, 13: 303, 26: 303, 43: 303, 51: 303, 303, 55: 303, 303, 59: 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303, 303},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 503},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 502},
		// 110
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 501},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 500},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 499},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 498},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 497},
		// 115
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 496},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 495},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 494},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 493},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 492},
		// 120
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 491},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 490},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 489},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 488},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 487},
		// 125
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 486},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 43: 480, 54: 478, 114: 479},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 477},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 476},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 475},
		// 130
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 474},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 473},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 472},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 471},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 470},
		// 135
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 469},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 468},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 467},
		{136, 10: 136, 136, 13: 136},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 261, 261, 261, 13: 430, 26: 261, 43: 261, 51: 261, 455, 55: 261, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		// 140
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 262, 262, 262, 13: 430, 26: 262, 43: 262, 51: 262, 455, 55: 262, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 263, 263, 263, 13: 430, 26: 263, 43: 263, 51: 263, 455, 55: 263, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 264, 264, 264, 13: 430, 26: 264, 43: 264, 51: 264, 455, 55: 264, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 265, 265, 265, 13: 430, 26: 265, 43: 265, 51: 265, 455, 55: 265, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 266, 266, 266, 13: 430, 26: 266, 43: 266, 51: 266, 455, 55: 266, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		// 145
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 267, 267, 267, 13: 430, 26: 267, 43: 267, 51: 267, 455, 55: 267, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 268, 268, 268, 13: 430, 26: 268, 43: 268, 51: 268, 455, 55: 268, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 269, 269, 269, 13: 430, 26: 269, 43: 269, 51: 269, 455, 55: 269, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 270, 270, 270, 13: 430, 26: 270, 43: 270, 51: 270, 455, 55: 270, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 271, 271, 271, 13: 430, 26: 271, 43: 271, 51: 271, 455, 55: 271, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		// 150
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 254, 254, 254, 13: 430, 43: 254, 52: 455, 55: 254, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{11: 483, 43: 482},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 481},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 257, 257, 257, 13: 430, 26: 257, 43: 257, 51: 257, 257, 55: 257, 432, 59: 438, 437, 443, 444, 454, 450, 451, 257, 257, 433, 257, 447, 446, 445, 441, 257, 257, 257, 448, 257, 453, 442, 257, 257, 257},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 485},
		// 155
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 484},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 253, 253, 253, 13: 430, 43: 253, 52: 455, 55: 253, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 272, 272, 272, 13: 430, 26: 272, 43: 272, 51: 272, 272, 55: 272, 432, 59: 438, 437, 443, 444, 454, 450, 451, 272, 272, 433, 272, 447, 446, 445, 441, 272, 272, 272, 448, 272, 453, 442, 272, 272, 272},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 273, 273, 273, 13: 430, 26: 273, 43: 273, 51: 273, 273, 55: 273, 432, 59: 438, 437, 443, 444, 273, 450, 451, 273, 273, 433, 273, 447, 446, 445, 441, 273, 273, 273, 448, 273, 273, 442, 273, 273, 273},
		{431, 436, 3: 449, 439, 440, 274, 435, 434, 274, 274, 274, 13: 430, 26: 274, 43: 274, 51: 274, 274, 55: 274, 432, 59: 438, 437, 443, 444, 274, 450, 451, 274, 274, 433, 274, 447, 446, 445, 441, 274, 274, 274, 448, 274, 274, 442, 274, 274, 274},
		// 160
		{431, 436, 3: 449, 439, 440, 275, 435, 434, 275, 275, 275, 13: 430, 26: 275, 43: 275, 51: 275, 275, 55: 275, 432, 59: 438, 437, 443, 444, 275, 450, 275, 275, 275, 433, 275, 447, 446, 445, 441, 275, 275, 275, 448, 275, 275, 442, 275, 275, 275},
		{431, 436, 3: 449, 439, 440, 276, 435, 434, 276, 276, 276, 13: 430, 26: 276, 43: 276, 51: 276, 276, 55: 276, 432, 59: 438, 437, 443, 444, 276, 276, 276, 276, 276, 433, 276, 447, 446, 445, 441, 276, 276, 276, 448, 276, 276, 442, 276, 276, 276},
		{431, 436, 3: 277, 439, 440, 277, 435, 434, 277, 277, 277, 13: 430, 26: 277, 43: 277, 51: 277, 277, 55: 277, 432, 59: 438, 437, 443, 444, 277, 277, 277, 277, 277, 433, 277, 447, 446, 445, 441, 277, 277, 277, 448, 277, 277, 442, 277, 277, 277},
		{431, 436, 3: 278, 439, 440, 278, 435, 434, 278, 278, 278, 13: 430, 26: 278, 43: 278, 51: 278, 278, 55: 278, 432, 59: 438, 437, 443, 444, 278, 278, 278, 278, 278, 433, 278, 278, 446, 445, 441, 278, 278, 278, 278, 278, 278, 442, 278, 278, 278},
		{431, 436, 3: 279, 439, 440, 279, 435, 434, 279, 279, 279, 13: 430, 26: 279, 43: 279, 51: 279, 279, 55: 279, 432, 59: 438, 437, 443, 444, 279, 279, 279, 279, 279, 433, 279, 279, 446, 445, 441, 279, 279, 279, 279, 279, 279, 442, 279, 279, 279},
		// 165
		{431, 436, 3: 280, 439, 440, 280, 435, 434, 280, 280, 280, 13: 430, 26: 280, 43: 280, 51: 280, 280, 55: 280, 432, 59: 438, 437, 280, 280, 280, 280, 280, 280, 280, 433, 280, 280, 280, 280, 441, 280, 280, 280, 280, 280, 280, 442, 280, 280, 280},
		{431, 436, 3: 281, 439, 440, 281, 435, 434, 281, 281, 281, 13: 430, 26: 281, 43: 281, 51: 281, 281, 55: 281, 432, 59: 438, 437, 281, 281, 281, 281, 281, 281, 281, 433, 281, 281, 281, 281, 441, 281, 281, 281, 281, 281, 281, 442, 281, 281, 281},
		{431, 436, 3: 282, 439, 440, 282, 435, 434, 282, 282, 282, 13: 430, 26: 282, 43: 282, 51: 282, 282, 55: 282, 432, 59: 438, 437, 282, 282, 282, 282, 282, 282, 282, 433, 282, 282, 282, 282, 441, 282, 282, 282, 282, 282, 282, 442, 282, 282, 282},
		{431, 436, 3: 283, 439, 440, 283, 435, 434, 283, 283, 283, 13: 430, 26: 283, 43: 283, 51: 283, 283, 55: 283, 432, 59: 438, 437, 283, 283, 283, 283, 283, 283, 283, 433, 283, 283, 283, 283, 441, 283, 283, 283, 283, 283, 283, 442, 283, 283, 283},
		{431, 436, 3: 284, 439, 440, 284, 435, 434, 284, 284, 284, 13: 430, 26: 284, 43: 284, 51: 284, 284, 55: 284, 432, 59: 438, 437, 284, 284, 284, 284, 284, 284, 284, 433, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284, 284},
		// 170
		{431, 436, 3: 285, 439, 440, 285, 435, 434, 285, 285, 285, 13: 430, 26: 285, 43: 285, 51: 285, 285, 55: 285, 432, 59: 438, 437, 285, 285, 285, 285, 285, 285, 285, 433, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285, 285},
		{431, 436, 3: 286, 286, 286, 286, 435, 434, 286, 286, 286, 13: 430, 26: 286, 43: 286, 51: 286, 286, 55: 286, 432, 59: 438, 437, 286, 286, 286, 286, 286, 286, 286, 433, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286, 286},
		{431, 436, 3: 287, 287, 287, 287, 435, 434, 287, 287, 287, 13: 430, 26: 287, 43: 287, 51: 287, 287, 55: 287, 432, 59: 438, 437, 287, 287, 287, 287, 287, 287, 287, 433, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287, 287},
		{431, 288, 3: 288, 288, 288, 288, 435, 434, 288, 288, 288, 13: 430, 26: 288, 43: 288, 51: 288, 288, 55: 288, 432, 59: 288, 288, 288, 288, 288, 288, 288, 288, 288, 433, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288, 288},
		{431, 289, 3: 289, 289, 289, 289, 435, 434, 289, 289, 289, 13: 430, 26: 289, 43: 289, 51: 289, 289, 55: 289, 432, 59: 289, 289, 289, 289, 289, 289, 289, 289, 289, 433, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289, 289},
		// 175
		{431, 290, 3: 290, 290, 290, 290, 435, 434, 290, 290, 290, 13: 430, 26: 290, 43: 290, 51: 290, 290, 55: 290, 432, 59: 290, 290, 290, 290, 290, 290, 290, 290, 290, 433, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290, 290},
		{305, 305, 3: 305, 305, 305, 305, 305, 305, 305, 305, 305, 13: 305, 26: 305, 43: 305, 51: 305, 305, 55: 305, 305, 59: 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305, 305},
		{306, 306, 3: 306, 306, 306, 306, 306, 306, 306, 306, 306, 13: 306, 26: 306, 43: 306, 51: 306, 306, 55: 306, 306, 59: 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306, 306},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 10: 320, 320, 13: 430, 52: 455, 56: 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{10: 317, 510},
		// 180
		{10: 509},
		{307, 307, 3: 307, 307, 307, 307, 307, 307, 307, 307, 307, 13: 307, 26: 307, 43: 307, 51: 307, 307, 55: 307, 307, 59: 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307, 307},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 511},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 10: 319, 319, 13: 430, 52: 455, 56: 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{11: 483, 55: 513},
		// 185
		{308, 308, 3: 308, 308, 308, 308, 308, 308, 308, 308, 308, 13: 308, 26: 308, 43: 308, 51: 308, 308, 55: 308, 308, 59: 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308, 308},
		{162, 162, 162, 162, 162, 162, 162, 162, 162, 10: 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 44: 162, 55: 162},
		{138, 10: 138, 138, 13: 138},
		{55: 520},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 518},
		// 190
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 13: 430, 52: 455, 55: 519, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{135, 10: 135, 135, 13: 135},
		{137, 10: 137, 137, 13: 137},
		{431, 298, 3: 298, 298, 298, 298, 435, 434, 298, 298, 298, 13: 430, 26: 298, 43: 298, 51: 298, 298, 55: 298, 432, 59: 298, 298, 298, 298, 298, 298, 298, 298, 298, 433, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298, 298},
		{134, 10: 134, 134, 13: 134},
		// 195
		{10: 566},
		{10: 157, 23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 113: 339, 117: 358, 361, 357, 338, 124: 528, 340, 127: 337, 161: 527, 172: 525, 526, 199: 529},
		{10: 159, 563},
		{10: 156},
		{10: 155, 155},
		// 200
		{165, 412, 165, 10: 141, 141, 13: 165, 133: 413, 136: 531, 532, 156: 416, 180: 533},
		{10: 530},
		{132, 10: 132, 132, 13: 132},
		{536, 2: 535, 13: 129, 185: 417, 419, 534},
		{10: 153, 153},
		// 205
		{10: 152, 152},
		{540, 9: 179, 179, 179, 13: 539, 23: 179, 179, 179, 27: 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 179, 52: 179, 179, 57: 179, 179},
		{176, 9: 176, 176, 176, 13: 176, 23: 176, 176, 176, 27: 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 52: 176, 176, 57: 176, 176},
		{165, 412, 165, 10: 133, 13: 165, 23: 133, 133, 133, 27: 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 44: 133, 133, 133, 133, 133, 133, 133, 133: 413, 136: 531, 537, 156: 523, 174: 524},
		{10: 538},
		// 210
		{175, 9: 175, 175, 175, 13: 175, 23: 175, 175, 175, 27: 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 52: 175, 175, 57: 175, 175},
		{161, 161, 161, 161, 161, 161, 161, 161, 161, 12: 161, 14: 161, 161, 161, 161, 161, 161, 161, 161, 161, 367, 365, 366, 44: 551, 55: 161, 113: 423, 154: 552, 163: 550},
		{2: 543, 10: 149, 23: 170, 170, 170, 27: 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 170, 44: 170, 170, 170, 170, 170, 170, 170, 170: 544, 194: 542, 213: 541},
		{23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 113: 339, 117: 358, 361, 357, 338, 124: 528, 340, 127: 337, 161: 527, 172: 525, 548},
		{10: 547},
		// 215
		{10: 151, 151, 147: 151},
		{10: 148, 545},
		{2: 546},
		{10: 150, 150, 147: 150},
		{168, 9: 168, 168, 168, 13: 168, 23: 168, 168, 168, 27: 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 168, 52: 168, 168, 57: 168, 168},
		// 220
		{10: 549},
		{169, 9: 169, 169, 169, 13: 169, 23: 169, 169, 169, 27: 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 169, 52: 169, 169, 57: 169, 169},
		{392, 559, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 421, 256, 169: 560},
		{161, 161, 161, 161, 161, 161, 161, 161, 161, 12: 161, 14: 161, 161, 161, 161, 161, 161, 161, 161, 161, 367, 365, 366, 113: 423, 154: 427, 163: 556},
		{160, 160, 160, 160, 160, 160, 160, 160, 160, 12: 160, 14: 160, 160, 160, 160, 160, 160, 160, 160, 160, 367, 365, 366, 44: 553, 55: 160, 113: 514},
		// 225
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 554},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 13: 430, 52: 455, 55: 555, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{172, 9: 172, 172, 172, 13: 172, 23: 172, 172, 172, 27: 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 172, 52: 172, 172, 57: 172, 172},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 557},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 13: 430, 52: 455, 55: 558, 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		// 230
		{173, 9: 173, 173, 173, 13: 173, 23: 173, 173, 173, 27: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 52: 173, 173, 57: 173, 173},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 521, 562},
		{55: 561},
		{174, 9: 174, 174, 174, 13: 174, 23: 174, 174, 174, 27: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 52: 174, 174, 57: 174, 174},
		{171, 9: 171, 171, 171, 13: 171, 23: 171, 171, 171, 27: 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 171, 52: 171, 171, 57: 171, 171},
		// 235
		{23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 113: 339, 117: 358, 361, 357, 338, 124: 528, 340, 127: 337, 147: 564, 161: 565},
		{10: 158},
		{10: 154, 154},
		{139, 10: 139, 139, 13: 139},
		{10: 131, 23: 131, 131, 131, 27: 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 44: 131, 131, 131, 131, 131, 131, 131, 205: 568},
		// 240
		{10: 157, 23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 113: 339, 117: 358, 361, 357, 338, 124: 528, 340, 127: 337, 161: 527, 172: 525, 526, 199: 569},
		{10: 570},
		{130, 10: 130, 130, 13: 130},
		{167, 412, 167, 10: 167, 167, 13: 167, 133: 572},
		{166, 2: 166, 10: 166, 166, 13: 166},
		// 245
		{200, 200, 200, 9: 200, 200, 13: 200, 43: 200},
		{198, 198, 198, 9: 198, 198, 13: 198, 43: 198},
		{201, 201, 201, 9: 201, 201, 13: 201, 43: 201},
		{260, 260, 3: 260, 260, 260, 260, 260, 260, 260, 260, 260, 13: 260, 26: 260, 43: 260, 51: 260, 260, 55: 260, 260, 59: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 145, 145, 145, 27: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 53: 582, 478, 114: 579, 123: 581, 155: 408, 162: 580},
		// 250
		{431, 293, 3: 293, 293, 293, 293, 435, 434, 293, 293, 293, 13: 430, 26: 293, 43: 293, 51: 293, 293, 55: 293, 432, 59: 293, 293, 293, 293, 293, 293, 293, 293, 293, 433, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293, 293},
		{10: 753, 483},
		{10: 747},
		{10: 746},
		{104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 12: 104, 14: 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 104, 44: 104, 104, 104, 104, 104, 104, 104, 53: 104, 57: 104, 104, 85: 104, 104, 104, 104, 104, 104, 104, 104, 104, 95: 104, 104, 206: 583},
		// 255
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 367, 365, 366, 100, 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 53: 582, 478, 57: 611, 377, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 113: 339, 585, 117: 358, 361, 357, 338, 601, 612, 588, 586, 340, 127: 337, 593, 589, 591, 592, 587, 134: 590, 600, 138: 336, 144: 599, 181: 597, 216: 598, 596},
		{316, 316, 3: 316, 316, 316, 316, 316, 316, 316, 11: 316, 13: 316, 43: 744, 52: 316, 56: 316, 59: 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316},
		{9: 251, 251, 483},
		{165, 412, 165, 9: 238, 133: 713, 136: 712, 711, 171: 709, 196: 710, 708},
		{114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 12: 114, 14: 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 114, 44: 114, 114, 114, 114, 114, 114, 114, 53: 114, 57: 114, 114, 85: 114, 114, 114, 114, 114, 114, 114, 114, 114, 95: 114, 114, 112: 114},
		// 260
		{113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 12: 113, 14: 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 113, 44: 113, 113, 113, 113, 113, 113, 113, 53: 113, 57: 113, 113, 85: 113, 113, 113, 113, 113, 113, 113, 113, 113, 95: 113, 113, 112: 113},
		{112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 12: 112, 14: 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 44: 112, 112, 112, 112, 112, 112, 112, 53: 112, 57: 112, 112, 85: 112, 112, 112, 112, 112, 112, 112, 112, 112, 95: 112, 112, 112: 112},
		{111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 12: 111, 14: 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 44: 111, 111, 111, 111, 111, 111, 111, 53: 111, 57: 111, 111, 85: 111, 111, 111, 111, 111, 111, 111, 111, 111, 95: 111, 111, 112: 111},
		{110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 12: 110, 14: 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 110, 44: 110, 110, 110, 110, 110, 110, 110, 53: 110, 57: 110, 110, 85: 110, 110, 110, 110, 110, 110, 110, 110, 110, 95: 110, 110, 112: 110},
		{109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 12: 109, 14: 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 109, 44: 109, 109, 109, 109, 109, 109, 109, 53: 109, 57: 109, 109, 85: 109, 109, 109, 109, 109, 109, 109, 109, 109, 95: 109, 109, 112: 109},
		// 265
		{108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 12: 108, 14: 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 108, 44: 108, 108, 108, 108, 108, 108, 108, 53: 108, 57: 108, 108, 85: 108, 108, 108, 108, 108, 108, 108, 108, 108, 95: 108, 108, 112: 108},
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 379, 705},
		{43: 703},
		{26: 702},
		{102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 12: 102, 14: 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 102, 44: 102, 102, 102, 102, 102, 102, 102, 53: 102, 57: 102, 102, 85: 102, 102, 102, 102, 102, 102, 102, 102, 102, 95: 102, 102},
		// 270
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 367, 365, 366, 99, 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 53: 582, 478, 57: 611, 377, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 113: 339, 585, 117: 358, 361, 357, 338, 601, 612, 588, 586, 340, 127: 337, 593, 589, 591, 592, 587, 134: 590, 600, 138: 336, 144: 599, 181: 701},
		{98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 12: 98, 14: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 44: 98, 98, 98, 98, 98, 98, 98, 53: 98, 57: 98, 98, 85: 98, 98, 98, 98, 98, 98, 98, 98, 98, 95: 98, 98},
		{97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 12: 97, 14: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 44: 97, 97, 97, 97, 97, 97, 97, 53: 97, 57: 97, 97, 85: 97, 97, 97, 97, 97, 97, 97, 97, 97, 95: 97, 97},
		{9: 700},
		{694},
		// 275
		{690},
		{686},
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 582, 478, 57: 611, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 114: 585, 121: 601, 612, 588, 128: 593, 589, 591, 592, 587, 134: 590, 680},
		{666},
		{392, 397, 662, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 663},
		// 280
		{9: 661},
		{9: 660},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 478, 114: 585, 121: 658},
		{60, 23: 614, 85: 60, 203: 613},
		{51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 12: 51, 14: 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 44: 51, 51, 51, 51, 51, 51, 51, 53: 51, 57: 51, 51, 85: 51, 51, 51, 51, 51, 51, 51, 51, 51, 95: 51, 51, 112: 51},
		// 285
		{615, 85: 616},
		{59, 85: 59},
		{12: 618, 164: 643},
		{617},
		{12: 618, 164: 619},
		// 290
		{10: 63, 12: 63, 43: 63},
		{12: 620, 43: 621},
		{10: 62, 12: 62, 43: 62},
		{43: 622},
		{12: 55, 626, 149: 624, 623, 157: 625},
		// 295
		{12: 639},
		{10: 57, 57, 43: 57},
		{11: 629, 43: 630},
		{2: 627},
		{55: 628},
		// 300
		{12: 54},
		{12: 55, 626, 149: 638, 623},
		{12: 631, 182: 632},
		{10: 53, 53, 43: 53},
		{11: 633, 43: 634},
		// 305
		{12: 637},
		{2: 543, 170: 635},
		{10: 636, 545},
		{47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 12: 47, 14: 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 44: 47, 47, 47, 47, 47, 47, 47, 53: 47, 57: 47, 47, 85: 47, 47, 47, 47, 47, 47, 47, 47, 47, 95: 47, 47, 112: 47},
		{10: 52, 52, 43: 52},
		// 310
		{10: 56, 56, 43: 56},
		{640},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 641},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 10: 642, 13: 430, 52: 455, 56: 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{10: 58, 58, 43: 58},
		// 315
		{10: 644, 12: 620, 43: 645},
		{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 12: 61, 14: 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 44: 61, 61, 61, 61, 61, 61, 61, 53: 61, 57: 61, 61, 85: 61, 61, 61, 61, 61, 61, 61, 61, 61, 95: 61, 61, 112: 61},
		{10: 647, 12: 55, 626, 43: 648, 149: 624, 623, 157: 646},
		{10: 651, 629, 43: 652},
		{46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 12: 46, 14: 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 44: 46, 46, 46, 46, 46, 46, 46, 53: 46, 57: 46, 46, 85: 46, 46, 46, 46, 46, 46, 46, 46, 46, 95: 46, 46, 112: 46},
		// 320
		{12: 55, 626, 149: 624, 623, 157: 649},
		{10: 650, 629},
		{45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 12: 45, 14: 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 44: 45, 45, 45, 45, 45, 45, 45, 53: 45, 57: 45, 45, 85: 45, 45, 45, 45, 45, 45, 45, 45, 45, 95: 45, 45, 112: 45},
		{50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 12: 50, 14: 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 44: 50, 50, 50, 50, 50, 50, 50, 53: 50, 57: 50, 50, 85: 50, 50, 50, 50, 50, 50, 50, 50, 50, 95: 50, 50, 112: 50},
		{12: 55, 626, 149: 624, 623, 157: 653},
		// 325
		{10: 654, 629, 43: 655},
		{49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 12: 49, 14: 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 49, 44: 49, 49, 49, 49, 49, 49, 49, 53: 49, 57: 49, 49, 85: 49, 49, 49, 49, 49, 49, 49, 49, 49, 95: 49, 49, 112: 49},
		{12: 631, 182: 656},
		{10: 657, 633},
		{48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 12: 48, 14: 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44: 48, 48, 48, 48, 48, 48, 48, 53: 48, 57: 48, 48, 85: 48, 48, 48, 48, 48, 48, 48, 48, 48, 95: 48, 48, 112: 48},
		// 330
		{9: 659},
		{85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 12: 85, 14: 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 44: 85, 85, 85, 85, 85, 85, 85, 53: 85, 57: 85, 85, 85: 85, 85, 85, 85, 85, 85, 85, 85, 85, 95: 85, 85, 112: 85},
		{86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 12: 86, 14: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 44: 86, 86, 86, 86, 86, 86, 86, 53: 86, 57: 86, 86, 85: 86, 86, 86, 86, 86, 86, 86, 86, 86, 95: 86, 86, 112: 86},
		{87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 12: 87, 14: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 44: 87, 87, 87, 87, 87, 87, 87, 53: 87, 57: 87, 87, 85: 87, 87, 87, 87, 87, 87, 87, 87, 87, 95: 87, 87, 112: 87},
		{316, 316, 3: 316, 316, 316, 316, 316, 316, 665, 13: 316, 52: 316, 56: 316, 59: 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316},
		// 335
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 664, 13: 430, 52: 455, 56: 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 12: 84, 14: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 44: 84, 84, 84, 84, 84, 84, 84, 53: 84, 57: 84, 84, 85: 84, 84, 84, 84, 84, 84, 84, 84, 84, 95: 84, 84, 112: 84},
		{88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 12: 88, 14: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 44: 88, 88, 88, 88, 88, 88, 88, 53: 88, 57: 88, 88, 85: 88, 88, 88, 88, 88, 88, 88, 88, 88, 95: 88, 88, 112: 88},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 54: 478, 58: 377, 113: 339, 585, 117: 358, 361, 357, 338, 667, 124: 586, 340, 127: 337, 138: 336, 144: 668},
		{9: 674},
		// 340
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 478, 114: 585, 121: 669},
		{9: 670},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 10: 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 478, 114: 585, 121: 671},
		{10: 672},
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 582, 478, 57: 611, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 114: 585, 121: 601, 612, 588, 128: 593, 589, 591, 592, 587, 134: 590, 673},
		// 345
		{89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 12: 89, 14: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 44: 89, 89, 89, 89, 89, 89, 89, 53: 89, 57: 89, 89, 85: 89, 89, 89, 89, 89, 89, 89, 89, 89, 95: 89, 89, 112: 89},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 478, 114: 585, 121: 675},
		{9: 676},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 10: 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 478, 114: 585, 121: 677},
		{10: 678},
		// 350
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 582, 478, 57: 611, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 114: 585, 121: 601, 612, 588, 128: 593, 589, 591, 592, 587, 134: 590, 679},
		{90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 12: 90, 14: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 44: 90, 90, 90, 90, 90, 90, 90, 53: 90, 57: 90, 90, 85: 90, 90, 90, 90, 90, 90, 90, 90, 90, 95: 90, 90, 112: 90},
		{86: 681},
		{682},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 478, 114: 683},
		// 355
		{10: 684, 483},
		{9: 685},
		{91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 12: 91, 14: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 44: 91, 91, 91, 91, 91, 91, 91, 53: 91, 57: 91, 91, 85: 91, 91, 91, 91, 91, 91, 91, 91, 91, 95: 91, 91, 112: 91},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 478, 114: 687},
		{10: 688, 483},
		// 360
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 582, 478, 57: 611, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 114: 585, 121: 601, 612, 588, 128: 593, 589, 591, 592, 587, 134: 590, 689},
		{92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 12: 92, 14: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 44: 92, 92, 92, 92, 92, 92, 92, 53: 92, 57: 92, 92, 85: 92, 92, 92, 92, 92, 92, 92, 92, 92, 95: 92, 92, 112: 92},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 478, 114: 691},
		{10: 692, 483},
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 582, 478, 57: 611, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 114: 585, 121: 601, 612, 588, 128: 593, 589, 591, 592, 587, 134: 590, 693},
		// 365
		{93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 12: 93, 14: 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 44: 93, 93, 93, 93, 93, 93, 93, 53: 93, 57: 93, 93, 85: 93, 93, 93, 93, 93, 93, 93, 93, 93, 95: 93, 93, 112: 93},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 54: 478, 114: 695},
		{10: 696, 483},
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 582, 478, 57: 611, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 114: 585, 121: 601, 612, 588, 128: 593, 589, 591, 592, 587, 134: 590, 697},
		{95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 12: 95, 14: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 44: 95, 95, 95, 95, 95, 95, 95, 53: 95, 57: 95, 95, 85: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95: 95, 95, 112: 698},
		// 370
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 582, 478, 57: 611, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 114: 585, 121: 601, 612, 588, 128: 593, 589, 591, 592, 587, 134: 590, 699},
		{94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 12: 94, 14: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 44: 94, 94, 94, 94, 94, 94, 94, 53: 94, 57: 94, 94, 85: 94, 94, 94, 94, 94, 94, 94, 94, 94, 95: 94, 94, 112: 94},
		{96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 12: 96, 14: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 44: 96, 96, 96, 96, 96, 96, 96, 53: 96, 57: 96, 96, 85: 96, 96, 96, 96, 96, 96, 96, 96, 96, 95: 96, 96, 112: 96},
		{101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 12: 101, 14: 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 101, 44: 101, 101, 101, 101, 101, 101, 101, 53: 101, 57: 101, 101, 85: 101, 101, 101, 101, 101, 101, 101, 101, 101, 95: 101, 101},
		{103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 12: 103, 14: 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 103, 44: 103, 103, 103, 103, 103, 103, 103, 103, 53: 103, 57: 103, 103, 85: 103, 103, 103, 103, 103, 103, 103, 103, 103, 95: 103, 103, 112: 103},
		// 375
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 582, 478, 57: 611, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 114: 585, 121: 601, 612, 588, 128: 593, 589, 591, 592, 587, 134: 590, 704},
		{105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 12: 105, 14: 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 105, 44: 105, 105, 105, 105, 105, 105, 105, 53: 105, 57: 105, 105, 85: 105, 105, 105, 105, 105, 105, 105, 105, 105, 95: 105, 105, 112: 105},
		{43: 706},
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 582, 478, 57: 611, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 114: 585, 121: 601, 612, 588, 128: 593, 589, 591, 592, 587, 134: 590, 707},
		{106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 12: 106, 14: 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 106, 44: 106, 106, 106, 106, 106, 106, 106, 53: 106, 57: 106, 106, 85: 106, 106, 106, 106, 106, 106, 106, 106, 106, 95: 106, 106, 112: 106},
		// 380
		{9: 743},
		{9: 240, 11: 240},
		{9: 237, 11: 741},
		{9: 236, 11: 236, 52: 235, 179: 715},
		{714, 2: 535, 187: 534},
		// 385
		{164, 2: 164},
		{165, 412, 165, 133: 713, 136: 712, 537},
		{52: 716},
		{392, 397, 717, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 720, 718, 160: 719},
		{316, 316, 3: 316, 316, 316, 316, 316, 316, 316, 11: 316, 13: 316, 26: 316, 43: 739, 52: 316, 56: 316, 59: 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316, 316},
		// 390
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 127, 11: 127, 13: 430, 26: 127, 52: 455, 56: 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{9: 234, 11: 234},
		{120, 120, 120, 120, 120, 120, 120, 120, 120, 11: 122, 120, 726, 120, 120, 120, 120, 120, 120, 120, 120, 120, 26: 122, 53: 120, 56: 727, 159: 725, 166: 724, 722, 723, 198: 721},
		{11: 734, 26: 193, 165: 735},
		{392, 397, 717, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 720, 718, 160: 733},
		// 395
		{13: 726, 52: 731, 56: 727, 159: 732},
		{119, 119, 119, 119, 119, 119, 119, 119, 119, 12: 119, 14: 119, 119, 119, 119, 119, 119, 119, 119, 119, 53: 119},
		{13: 118, 52: 118, 56: 118},
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 379, 729},
		{2: 728},
		// 400
		{13: 115, 52: 115, 56: 115},
		{55: 730},
		{13: 116, 52: 116, 56: 116},
		{121, 121, 121, 121, 121, 121, 121, 121, 121, 12: 121, 14: 121, 121, 121, 121, 121, 121, 121, 121, 121, 53: 121},
		{13: 117, 52: 117, 56: 117},
		// 405
		{11: 124, 26: 124},
		{120, 120, 120, 120, 120, 120, 120, 120, 120, 12: 120, 726, 120, 120, 120, 120, 120, 120, 120, 120, 120, 26: 192, 53: 120, 56: 727, 159: 725, 166: 724, 737, 723},
		{26: 736},
		{9: 126, 11: 126, 26: 126},
		{392, 397, 717, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 720, 718, 160: 738},
		// 410
		{11: 123, 26: 123},
		{392, 397, 717, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 720, 718, 160: 740},
		{9: 125, 11: 125, 26: 125},
		{165, 412, 165, 133: 713, 136: 712, 711, 171: 742},
		{9: 239, 11: 239},
		// 415
		{248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 12: 248, 14: 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 44: 248, 248, 248, 248, 248, 248, 248, 248, 53: 248, 57: 248, 248, 85: 248, 248, 248, 248, 248, 248, 248, 248, 248, 95: 248, 248},
		{392, 397, 584, 396, 398, 399, 404, 395, 394, 252, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 582, 478, 57: 611, 85: 607, 604, 609, 594, 608, 595, 605, 606, 602, 95: 610, 603, 114: 585, 121: 601, 612, 588, 128: 593, 589, 591, 592, 587, 134: 590, 745},
		{107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 12: 107, 14: 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 107, 44: 107, 107, 107, 107, 107, 107, 107, 53: 107, 57: 107, 107, 85: 107, 107, 107, 107, 107, 107, 107, 107, 107, 95: 107, 107, 112: 107},
		{259, 259, 3: 259, 259, 259, 259, 259, 259, 259, 259, 259, 13: 259, 26: 259, 43: 259, 51: 259, 259, 55: 259, 259, 59: 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259},
		{392, 292, 385, 292, 292, 292, 292, 395, 394, 292, 292, 292, 391, 292, 401, 400, 403, 386, 387, 388, 389, 390, 402, 26: 292, 43: 292, 51: 292, 292, 749, 748, 292, 292, 59: 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292, 292},
		// 420
		{431, 291, 3: 291, 291, 291, 291, 435, 434, 291, 291, 291, 13: 430, 26: 291, 43: 291, 51: 291, 291, 55: 291, 432, 59: 291, 291, 291, 291, 291, 291, 291, 291, 291, 433, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291, 291},
		{120, 120, 120, 120, 120, 120, 120, 120, 120, 11: 122, 120, 726, 120, 120, 120, 120, 120, 120, 120, 120, 120, 26: 122, 53: 120, 56: 727, 159: 725, 166: 724, 722, 723, 198: 750},
		{11: 734, 26: 193, 165: 751},
		{26: 752},
		{302, 302, 3: 302, 302, 302, 302, 302, 302, 302, 302, 302, 13: 302, 26: 302, 43: 302, 51: 302, 302, 55: 302, 302, 59: 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302, 302},
		// 425
		{309, 309, 3: 309, 309, 309, 309, 309, 309, 309, 309, 309, 13: 309, 26: 309, 43: 309, 51: 309, 309, 55: 309, 309, 59: 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309, 309},
		{431, 294, 3: 294, 294, 294, 294, 435, 434, 294, 294, 294, 13: 430, 26: 294, 43: 294, 51: 294, 294, 55: 294, 432, 59: 294, 294, 294, 294, 294, 294, 294, 294, 294, 433, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294, 294},
		{431, 295, 3: 295, 295, 295, 295, 435, 434, 295, 295, 295, 13: 430, 26: 295, 43: 295, 51: 295, 295, 55: 295, 432, 59: 295, 295, 295, 295, 295, 295, 295, 295, 295, 433, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295, 295},
		{431, 296, 3: 296, 296, 296, 296, 435, 434, 296, 296, 296, 13: 430, 26: 296, 43: 296, 51: 296, 296, 55: 296, 432, 59: 296, 296, 296, 296, 296, 296, 296, 296, 296, 433, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296, 296},
		{431, 297, 3: 297, 297, 297, 297, 435, 434, 297, 297, 297, 13: 430, 26: 297, 43: 297, 51: 297, 297, 55: 297, 432, 59: 297, 297, 297, 297, 297, 297, 297, 297, 297, 433, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297, 297},
		// 430
		{431, 299, 3: 299, 299, 299, 299, 435, 434, 299, 299, 299, 13: 430, 26: 299, 43: 299, 51: 299, 299, 55: 299, 432, 59: 299, 299, 299, 299, 299, 299, 299, 299, 299, 433, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299, 299},
		{431, 300, 3: 300, 300, 300, 300, 435, 434, 300, 300, 300, 13: 430, 26: 300, 43: 300, 51: 300, 300, 55: 300, 432, 59: 300, 300, 300, 300, 300, 300, 300, 300, 300, 433, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300, 300},
		{431, 301, 3: 301, 301, 301, 301, 435, 434, 301, 301, 301, 13: 430, 26: 301, 43: 301, 51: 301, 301, 55: 301, 432, 59: 301, 301, 301, 301, 301, 301, 301, 301, 301, 433, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301, 301},
		{10: 762},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 53: 749, 748},
		// 435
		{764},
		{12: 618, 164: 765},
		{10: 644, 12: 620},
		{23: 65, 65, 65, 27: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 44: 65, 65, 65, 65, 65, 65, 65, 53: 66, 57: 66, 65, 178: 768, 183: 767},
		{53: 74, 57: 74, 208: 772},
		// 440
		{23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 58: 377, 113: 339, 117: 358, 361, 357, 338, 124: 586, 340, 127: 337, 138: 336, 144: 769, 219: 770},
		{23: 68, 68, 68, 27: 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 44: 68, 68, 68, 68, 68, 68, 68, 53: 68, 57: 68, 68},
		{23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 44: 343, 344, 342, 368, 369, 345, 341, 53: 64, 57: 64, 377, 113: 339, 117: 358, 361, 357, 338, 124: 586, 340, 127: 337, 138: 336, 144: 771},
		{23: 67, 67, 67, 27: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 44: 67, 67, 67, 67, 67, 67, 67, 53: 67, 57: 67, 67},
		{53: 72, 57: 70, 176: 774, 775, 192: 773},
		// 445
		{73, 73, 73, 9: 73, 23: 73, 73, 73, 27: 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 73, 44: 73, 73, 73, 73, 73, 73, 73, 73, 57: 73, 73},
		{53: 582, 123: 778},
		{57: 611, 122: 612, 128: 776},
		{9: 777},
		{69, 69, 69, 9: 69, 23: 69, 69, 69, 27: 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 69, 44: 69, 69, 69, 69, 69, 69, 69, 69, 57: 69, 69},
		// 450
		{71, 71, 71, 9: 71, 23: 71, 71, 71, 27: 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 71, 44: 71, 71, 71, 71, 71, 71, 71, 71, 57: 71, 71},
		{79, 79, 79, 9: 79, 23: 79, 79, 79, 27: 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 79, 44: 79, 79, 79, 79, 79, 79, 79, 79, 57: 79, 79},
		{53: 191, 212: 782},
		{189, 189, 189, 9: 189, 189, 189, 13: 189, 23: 189, 189, 189, 27: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 53: 146},
		{53: 783},
		// 455
		{2: 784, 189: 787, 786, 225: 785},
		{11: 321, 26: 321, 52: 321},
		{11: 790, 26: 193, 165: 791},
		{11: 188, 26: 188},
		{11: 186, 26: 186, 52: 788},
		// 460
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 379, 789},
		{11: 185, 26: 185},
		{2: 784, 26: 192, 189: 787, 793},
		{26: 792},
		{190, 190, 190, 9: 190, 190, 190, 13: 190, 23: 190, 190, 190, 27: 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190, 190},
		// 465
		{11: 187, 26: 187},
		{53: 796},
		{210, 210, 210, 9: 210, 210, 210, 13: 210, 23: 210, 210, 210, 27: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 53: 146},
		{23: 212, 212, 212, 798, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 58: 212, 211: 797},
		{23: 367, 365, 366, 27: 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 58: 377, 113: 410, 117: 358, 361, 357, 409, 138: 802, 153: 801, 201: 800, 230: 799},
		// 470
		{209, 209, 209, 9: 209, 209, 209, 13: 209, 23: 209, 209, 209, 27: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209},
		{23: 367, 365, 366, 813, 355, 347, 356, 352, 364, 351, 349, 350, 348, 353, 362, 359, 360, 363, 354, 346, 58: 377, 113: 410, 117: 358, 361, 357, 409, 138: 802, 153: 801, 201: 814},
		{23: 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 58: 206},
		{165, 412, 165, 9: 804, 43: 178, 133: 713, 136: 712, 806, 184: 807, 202: 805, 231: 803},
		{23: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 58: 202},
		// 475
		{9: 810, 11: 811},
		{23: 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 58: 203},
		{9: 197, 11: 197},
		{9: 195, 11: 195, 43: 177},
		{43: 808},
		// 480
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 379, 809},
		{9: 194, 11: 194},
		{23: 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 58: 204},
		{165, 412, 165, 43: 178, 133: 713, 136: 712, 806, 184: 807, 202: 812},
		{9: 196, 11: 196},
		// 485
		{211, 211, 211, 9: 211, 211, 211, 13: 211, 23: 211, 211, 211, 27: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211},
		{23: 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 205, 58: 205},
		{392, 397, 385, 396, 398, 399, 404, 395, 394, 12: 391, 14: 401, 400, 403, 386, 387, 388, 389, 390, 402, 145, 145, 145, 27: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 54: 816, 155: 408, 162: 817},
		{431, 436, 3: 449, 439, 440, 452, 435, 434, 10: 819, 13: 430, 52: 455, 56: 432, 59: 438, 437, 443, 444, 454, 450, 451, 459, 463, 433, 457, 447, 446, 445, 441, 461, 458, 456, 448, 465, 453, 442, 462, 460, 464},
		{10: 818},
		// 490
		{213, 213, 213, 9: 213, 213, 213, 13: 213, 23: 213, 213, 213, 27: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213},
		{214, 214, 214, 9: 214, 214, 214, 13: 214, 23: 214, 214, 214, 27: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214},
		{243, 243, 243, 9: 243, 243, 243, 13: 243},
		{241, 241, 241, 9: 241, 241, 241, 13: 241},
		{244, 244, 244, 9: 244, 244, 244, 13: 244},
		// 495
		{245, 245, 245, 9: 245, 245, 245, 13: 245},
		{246, 246, 246, 9: 246, 246, 246, 13: 246},
		{9: 236, 11: 236, 23: 65, 65, 65, 27: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 44: 65, 65, 65, 65, 65, 65, 65, 52: 235, 66, 57: 66, 65, 178: 768, 715, 183: 826},
		{53: 77, 57: 77, 207: 827},
		{53: 72, 57: 70, 176: 774, 775, 192: 828},
		// 500
		{76, 76, 76, 9: 76, 23: 76, 76, 76, 27: 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 44: 76, 76, 76, 76, 76, 76, 76, 76, 57: 76, 76},
		{82, 82, 82, 9: 82, 23: 82, 82, 82, 27: 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 82, 44: 82, 82, 82, 82, 82, 82, 82, 82, 57: 82, 82},
		{250, 250, 250, 250, 250, 250, 250, 250, 250, 12: 250, 14: 250, 250, 250, 250, 250, 250, 250, 250, 250, 145: 379, 831},
		{51: 324},
		{84: 854, 94: 856, 100: 844, 845, 846, 841, 842, 843, 847, 851, 848, 838, 849, 850, 115: 855, 853, 126: 852, 139: 836, 835, 840, 837, 839, 148: 834, 228: 833},
		// 505
		{51: 326},
		{51: 43, 84: 854, 94: 856, 100: 844, 845, 846, 841, 842, 843, 847, 851, 848, 838, 849, 850, 115: 855, 853, 126: 852, 139: 836, 894, 840, 837, 839},
		{51: 42, 84: 42, 94: 42, 97: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42},
		{51: 38, 84: 38, 94: 38, 97: 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{51: 37, 84: 37, 94: 37, 97: 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		// 510
		{94: 856, 115: 916, 853},
		{51: 35, 84: 35, 94: 35, 97: 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		{97: 28, 28, 904, 188: 902, 220: 903, 901},
		{94: 856, 115: 898, 853},
		{2: 895},
		// 515
		{2: 890},
		{2: 871, 84: 873, 226: 872},
		{84: 854, 94: 856, 115: 855, 853, 126: 870},
		{51: 16, 84: 16, 94: 16, 97: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{94: 856, 115: 868, 853},
		// 520
		{94: 856, 115: 866, 853},
		{84: 854, 94: 856, 115: 855, 853, 126: 865},
		{2: 861},
		{94: 856, 115: 859, 853},
		{51: 7, 84: 7, 94: 7, 97: 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		// 525
		{84: 5, 94: 858},
		{51: 4, 84: 4, 94: 4, 97: 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{84: 857},
		{84: 2, 94: 2},
		{51: 3, 84: 3, 94: 3, 97: 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		// 530
		{84: 1, 94: 1},
		{84: 860},
		{51: 8, 84: 8, 94: 8, 97: 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{84: 862, 94: 856, 115: 863, 853},
		{51: 12, 84: 12, 94: 12, 97: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		// 535
		{84: 864},
		{51: 9, 84: 9, 94: 9, 97: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{51: 13, 84: 13, 94: 13, 97: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{84: 867},
		{51: 14, 84: 14, 94: 14, 97: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		// 540
		{84: 869},
		{51: 15, 84: 15, 94: 15, 97: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{51: 17, 84: 17, 94: 17, 97: 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{84: 854, 94: 856, 115: 855, 853, 126: 879, 152: 889},
		{2: 543, 10: 149, 147: 875, 170: 874, 194: 876},
		// 545
		{51: 10, 84: 10, 94: 10, 97: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{10: 148, 882, 147: 883},
		{10: 880},
		{10: 877},
		{84: 854, 94: 856, 115: 855, 853, 126: 879, 152: 878},
		// 550
		{51: 18, 84: 18, 94: 18, 97: 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{51: 6, 84: 6, 94: 6, 97: 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{84: 854, 94: 856, 115: 855, 853, 126: 879, 152: 881},
		{51: 20, 84: 20, 94: 20, 97: 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{2: 546, 147: 886},
		// 555
		{10: 884},
		{84: 854, 94: 856, 115: 855, 853, 126: 879, 152: 885},
		{51: 11, 84: 11, 94: 11, 97: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{10: 887},
		{84: 854, 94: 856, 115: 855, 853, 126: 879, 152: 888},
		// 560
		{51: 19, 84: 19, 94: 19, 97: 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{51: 21, 84: 21, 94: 21, 97: 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{84: 891},
		{84: 854, 94: 856, 97: 40, 40, 40, 844, 845, 846, 841, 842, 843, 847, 851, 848, 838, 849, 850, 115: 855, 853, 126: 852, 139: 836, 835, 840, 837, 839, 148: 892, 151: 893},
		{84: 854, 94: 856, 97: 39, 39, 39, 844, 845, 846, 841, 842, 843, 847, 851, 848, 838, 849, 850, 115: 855, 853, 126: 852, 139: 836, 894, 840, 837, 839},
		// 565
		{97: 31, 31, 31},
		{51: 41, 84: 41, 94: 41, 97: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41},
		{84: 896},
		{84: 854, 94: 856, 97: 40, 40, 40, 844, 845, 846, 841, 842, 843, 847, 851, 848, 838, 849, 850, 115: 855, 853, 126: 852, 139: 836, 835, 840, 837, 839, 148: 892, 151: 897},
		{97: 32, 32, 32},
		// 570
		{84: 899},
		{84: 854, 94: 856, 97: 40, 40, 40, 844, 845, 846, 841, 842, 843, 847, 851, 848, 838, 849, 850, 115: 855, 853, 126: 852, 139: 836, 835, 840, 837, 839, 148: 892, 151: 900},
		{97: 33, 33, 33},
		{97: 24, 910, 222: 911, 909},
		{97: 30, 30, 30},
		// 575
		{97: 27, 27, 904, 188: 908},
		{94: 856, 115: 905, 853},
		{84: 906},
		{84: 854, 94: 856, 97: 40, 40, 40, 844, 845, 846, 841, 842, 843, 847, 851, 848, 838, 849, 850, 115: 855, 853, 126: 852, 139: 836, 835, 840, 837, 839, 148: 892, 151: 907},
		{97: 26, 26, 26},
		// 580
		{97: 29, 29, 29},
		{97: 915, 224: 914},
		{84: 912},
		{97: 23},
		{84: 854, 94: 856, 97: 40, 100: 844, 845, 846, 841, 842, 843, 847, 851, 848, 838, 849, 850, 115: 855, 853, 126: 852, 139: 836, 835, 840, 837, 839, 148: 892, 151: 913},
		// 585
		{97: 25},
		{51: 34, 84: 34, 94: 34, 97: 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{51: 22, 84: 22, 94: 22, 97: 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{84: 917},
		{51: 36, 84: 36, 94: 36, 97: 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("'%c'", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), PrettyString(lval.Token): %v\n", yySymName(n), n, n, PrettyString(lval.Token))
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 236

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			lx := yylex.(*lexer)
			lx.preprocessingFile = nil
		}
	case 2:
		{
			lx := yylex.(*lexer)
			lx.preprocessingFile = yyS[yypt-0].node.(*PreprocessingFile)
		}
	case 3:
		{
			lx := yylex.(*lexer)
			lx.constantExpression = nil
		}
	case 4:
		{
			lx := yylex.(*lexer)
			lx.constantExpression = yyS[yypt-0].node.(*ConstantExpression)
		}
	case 5:
		{
			lx := yylex.(*lexer)
			lx.translationUnit = nil
		}
	case 6:
		{
			lx := yylex.(*lexer)
			if lx.report.Errors(false) == nil && lx.scope.kind != ScopeFile {
				panic("internal error")
			}

			lx.translationUnit = yyS[yypt-0].node.(*TranslationUnit).reverse()
			lx.translationUnit.Declarations = lx.scope
		}
	case 7:
		{
			yyVAL.node = &EnumerationConstant{
				Token: yyS[yypt-0].Token,
			}
		}
	case 8:
		{
			yyVAL.node = &ArgumentExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 9:
		{
			yyVAL.node = &ArgumentExpressionList{
				Case: 1,
				ArgumentExpressionList: yyS[yypt-2].node.(*ArgumentExpressionList),
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 10:
		{
			yyVAL.node = (*ArgumentExpressionListOpt)(nil)
		}
	case 11:
		{
			yyVAL.node = &ArgumentExpressionListOpt{
				ArgumentExpressionList: yyS[yypt-0].node.(*ArgumentExpressionList).reverse(),
			}
		}
	case 12:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.scope = lx.scope
		}
	case 13:
		{
			yyVAL.node = &Expression{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 14:
		{
			yyVAL.node = &Expression{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
		}
	case 15:
		{
			yyVAL.node = &Expression{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 16:
		{
			yyVAL.node = &Expression{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
		}
	case 17:
		{
			yyVAL.node = &Expression{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 18:
		{
			yyVAL.node = &Expression{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
		}
	case 19:
		{
			yyVAL.node = &Expression{
				Case:           7,
				Token:          yyS[yypt-2].Token,
				ExpressionList: yyS[yypt-1].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 20:
		{
			yyVAL.node = &Expression{
				Case:           8,
				Expression:     yyS[yypt-3].node.(*Expression),
				Token:          yyS[yypt-2].Token,
				ExpressionList: yyS[yypt-1].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 21:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Case:       9,
				Expression: yyS[yypt-3].node.(*Expression),
				Token:      yyS[yypt-2].Token,
				ArgumentExpressionListOpt: yyS[yypt-1].node.(*ArgumentExpressionListOpt),
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			o := lhs.ArgumentExpressionListOpt
			if o == nil {
				break
			}

			if lhs.Expression.Case == 0 { // IDENTIFIER
				if lx.tweaks.enableBuiltinConstantP && lhs.Expression.Token.Val == idBuiltinConstantP {
					break
				}

				b := lhs.Expression.scope.Lookup(NSIdentifiers, lhs.Expression.Token.Val)
				if b.Node == nil && lx.tweaks.enableImplicitFuncDef {
					for l := o.ArgumentExpressionList; l != nil; l = l.ArgumentExpressionList {
						l.Expression.eval(lx)
					}
					break
				}
			}

			lhs.Expression.eval(lx)
			for l := o.ArgumentExpressionList; l != nil; l = l.ArgumentExpressionList {
				l.Expression.eval(lx)
			}
		}
	case 22:
		{
			yyVAL.node = &Expression{
				Case:       10,
				Expression: yyS[yypt-2].node.(*Expression),
				Token:      yyS[yypt-1].Token,
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 23:
		{
			yyVAL.node = &Expression{
				Case:       11,
				Expression: yyS[yypt-2].node.(*Expression),
				Token:      yyS[yypt-1].Token,
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 24:
		{
			yyVAL.node = &Expression{
				Case:       12,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 25:
		{
			yyVAL.node = &Expression{
				Case:       13,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 26:
		{
			yyVAL.node = &Expression{
				Case:            14,
				Token:           yyS[yypt-6].Token,
				TypeName:        yyS[yypt-5].node.(*TypeName),
				Token2:          yyS[yypt-4].Token,
				Token3:          yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token4:          yyS[yypt-0].Token,
			}
		}
	case 27:
		{
			yyVAL.node = &Expression{
				Case:       15,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 28:
		{
			yyVAL.node = &Expression{
				Case:       16,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 29:
		{
			yyVAL.node = &Expression{
				Case:       17,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 30:
		{
			yyVAL.node = &Expression{
				Case:       18,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 31:
		{
			yyVAL.node = &Expression{
				Case:       19,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 32:
		{
			yyVAL.node = &Expression{
				Case:       20,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 33:
		{
			yyVAL.node = &Expression{
				Case:       21,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 34:
		{
			yyVAL.node = &Expression{
				Case:       22,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 35:
		{
			yyVAL.node = &Expression{
				Case:       23,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 36:
		{
			yyVAL.node = &Expression{
				Case:     24,
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				TypeName: yyS[yypt-1].node.(*TypeName),
				Token3:   yyS[yypt-0].Token,
			}
		}
	case 37:
		{
			yyVAL.node = &Expression{
				Case:       25,
				Token:      yyS[yypt-3].Token,
				TypeName:   yyS[yypt-2].node.(*TypeName),
				Token2:     yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 38:
		{
			yyVAL.node = &Expression{
				Case:        26,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 39:
		{
			yyVAL.node = &Expression{
				Case:        27,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 40:
		{
			yyVAL.node = &Expression{
				Case:        28,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 41:
		{
			yyVAL.node = &Expression{
				Case:        29,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 42:
		{
			yyVAL.node = &Expression{
				Case:        30,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 43:
		{
			yyVAL.node = &Expression{
				Case:        31,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 44:
		{
			yyVAL.node = &Expression{
				Case:        32,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 45:
		{
			yyVAL.node = &Expression{
				Case:        33,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 46:
		{
			yyVAL.node = &Expression{
				Case:        34,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 47:
		{
			yyVAL.node = &Expression{
				Case:        35,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 48:
		{
			yyVAL.node = &Expression{
				Case:        36,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 49:
		{
			yyVAL.node = &Expression{
				Case:        37,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 50:
		{
			yyVAL.node = &Expression{
				Case:        38,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 51:
		{
			yyVAL.node = &Expression{
				Case:        39,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 52:
		{
			yyVAL.node = &Expression{
				Case:        40,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 53:
		{
			yyVAL.node = &Expression{
				Case:        41,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 54:
		{
			yyVAL.node = &Expression{
				Case:        42,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 55:
		{
			yyVAL.node = &Expression{
				Case:        43,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 56:
		{
			yyVAL.node = &Expression{
				Case:           44,
				Expression:     yyS[yypt-4].node.(*Expression),
				Token:          yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token2:         yyS[yypt-1].Token,
				Expression2:    yyS[yypt-0].node.(*Expression),
			}
		}
	case 57:
		{
			yyVAL.node = &Expression{
				Case:        45,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 58:
		{
			yyVAL.node = &Expression{
				Case:        46,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 59:
		{
			yyVAL.node = &Expression{
				Case:        47,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 60:
		{
			yyVAL.node = &Expression{
				Case:        48,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 61:
		{
			yyVAL.node = &Expression{
				Case:        49,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 62:
		{
			yyVAL.node = &Expression{
				Case:        50,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 63:
		{
			yyVAL.node = &Expression{
				Case:        51,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 64:
		{
			yyVAL.node = &Expression{
				Case:        52,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 65:
		{
			yyVAL.node = &Expression{
				Case:        53,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 66:
		{
			yyVAL.node = &Expression{
				Case:        54,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 67:
		{
			yyVAL.node = &Expression{
				Case:        55,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 68:
		{
			yyVAL.node = &Expression{
				Case:     56,
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				TypeName: yyS[yypt-1].node.(*TypeName),
				Token3:   yyS[yypt-0].Token,
			}
		}
	case 69:
		{
			yyVAL.node = &Expression{
				Case:              57,
				Token:             yyS[yypt-2].Token,
				CompoundStatement: yyS[yypt-1].node.(*CompoundStatement),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 70:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Case:   58,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableComputedGotos {
				lx.report.Err(lhs.Pos(), "computed gotos not enabled")
			}
		}
	case 71:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Case:        59,
				Expression:  yyS[yypt-3].node.(*Expression),
				Token:       yyS[yypt-2].Token,
				Token2:      yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableOmitConditionalOperand {
				lx.report.Err(lhs.Pos(), "omitting conditional operand not enabled")
			}
		}
	case 72:
		{
			yyVAL.node = (*ExpressionOpt)(nil)
		}
	case 73:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionOpt{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 74:
		{
			yyVAL.node = &ExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 75:
		{
			yyVAL.node = &ExpressionList{
				Case:           1,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList),
				Token:          yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 76:
		{
			yyVAL.node = (*ExpressionListOpt)(nil)
		}
	case 77:
		{
			lx := yylex.(*lexer)
			lhs := &ExpressionListOpt{
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 78:
		{
			lx := yylex.(*lexer)
			lx.constExprToks = []xc.Token{lx.last}
		}
	case 79:
		{
			lx := yylex.(*lexer)
			lhs := &ConstantExpression{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Value, lhs.Type = lhs.Expression.eval(lx)
			if lhs.Value == nil {
				lx.report.Err(lhs.Pos(), "not a constant expression")
			}
			l := lx.constExprToks
			lhs.toks = l[:len(l)-1]
			lx.constExprToks = nil
		}
	case 80:
		{
			lx := yylex.(*lexer)
			lhs := &Declaration{
				DeclarationSpecifiers: yyS[yypt-2].node.(*DeclarationSpecifiers),
				InitDeclaratorListOpt: yyS[yypt-1].node.(*InitDeclaratorListOpt),
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			ts0 := lhs.DeclarationSpecifiers.typeSpecifiers()
			if ts0 == 0 && lx.tweaks.enableImplicitIntType {
				lhs.DeclarationSpecifiers.typeSpecifier = tsEncode(tsInt)
			}
			ts := tsDecode(lhs.DeclarationSpecifiers.typeSpecifiers())
			ok := false
			for _, v := range ts {
				if v == tsStructSpecifier || v == tsUnionSpecifier {
					ok = true
					break
				}
			}
			if ok {
				s := lhs.DeclarationSpecifiers
				d := &Declarator{specifier: s}
				dd := &DirectDeclarator{
					Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
					declarator: d,
					idScope:    lx.scope,
					specifier:  s,
				}
				d.DirectDeclarator = dd
				d.setFull(lx)
				for l := lhs.DeclarationSpecifiers; l != nil; {
					ts := l.TypeSpecifier
					if ts != nil && ts.Case == 11 && ts.StructOrUnionSpecifier.Case == 0 { // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
						ts.StructOrUnionSpecifier.declarator = d
						break
					}

					if o := l.DeclarationSpecifiersOpt; o != nil {
						l = o.DeclarationSpecifiers
						continue
					}

					break
				}
			}

			o := lhs.InitDeclaratorListOpt
			if o != nil {
				break
			}

			s := lhs.DeclarationSpecifiers
			d := &Declarator{specifier: s}
			dd := &DirectDeclarator{
				Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
				declarator: d,
				idScope:    lx.scope,
				specifier:  s,
			}
			d.DirectDeclarator = dd
			d.setFull(lx)
			lhs.declarator = d
		}
	case 81:
		{
			yyVAL.node = &Declaration{
				Case: 1,
				StaticAssertDeclaration: yyS[yypt-0].node.(*StaticAssertDeclaration),
			}
		}
	case 82:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				StorageClassSpecifier:    yyS[yypt-1].node.(*StorageClassSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.StorageClassSpecifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid storage class specifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.StorageClassSpecifier.Case != 0 /* "typedef" */ && lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid storage class specifier")
			}
		}
	case 83:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     1,
				TypeSpecifier:            yyS[yypt-1].node.(*TypeSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeSpecifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.typeSpecifier = a.typeSpecifier
				break
			}

			lhs.attr = b.attr
			tsb := tsDecode(b.typeSpecifier)
			if len(tsb) == 1 && tsb[0] == tsTypedefName && lx.tweaks.allowCompatibleTypedefRedefinitions {
				tsb[0] = 0
			}
			ts := tsEncode(append(tsDecode(a.typeSpecifier), tsb...)...)
			if _, ok := tsValid[ts]; !ok {
				ts = tsEncode(tsInt)
				lx.report.Err(a.Pos(), "invalid type specifier")
			}
			lhs.typeSpecifier = ts
		}
	case 84:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     2,
				TypeQualifier:            yyS[yypt-1].node.(*TypeQualifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeQualifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid type qualifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid type qualifier")
			}
		}
	case 85:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationSpecifiers{
				Case:                     3,
				FunctionSpecifier:        yyS[yypt-1].node.(*FunctionSpecifier),
				DeclarationSpecifiersOpt: yyS[yypt-0].node.(*DeclarationSpecifiersOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.FunctionSpecifier
			b := lhs.DeclarationSpecifiersOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid function specifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
			if lhs.IsTypedef() {
				lx.report.Err(a.Pos(), "invalid function specifier")
			}
		}
	case 86:
		{
			yyVAL.node = (*DeclarationSpecifiersOpt)(nil)
		}
	case 87:
		{
			lhs := &DeclarationSpecifiersOpt{
				DeclarationSpecifiers: yyS[yypt-0].node.(*DeclarationSpecifiers),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.DeclarationSpecifiers.attr
			lhs.typeSpecifier = lhs.DeclarationSpecifiers.typeSpecifier
		}
	case 88:
		{
			yyVAL.node = &InitDeclaratorList{
				InitDeclarator: yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 89:
		{
			yyVAL.node = &InitDeclaratorList{
				Case:               1,
				InitDeclaratorList: yyS[yypt-2].node.(*InitDeclaratorList),
				Token:              yyS[yypt-1].Token,
				InitDeclarator:     yyS[yypt-0].node.(*InitDeclarator),
			}
		}
	case 90:
		{
			yyVAL.node = (*InitDeclaratorListOpt)(nil)
		}
	case 91:
		{
			yyVAL.node = &InitDeclaratorListOpt{
				InitDeclaratorList: yyS[yypt-0].node.(*InitDeclaratorList).reverse(),
			}
		}
	case 92:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
		}
	case 93:
		{
			lx := yylex.(*lexer)
			d := yyS[yypt-0].node.(*Declarator)
			d.setFull(lx)
		}
	case 94:
		{
			lx := yylex.(*lexer)
			lhs := &InitDeclarator{
				Case:        1,
				Declarator:  yyS[yypt-3].node.(*Declarator),
				Token:       yyS[yypt-1].Token,
				Initializer: yyS[yypt-0].node.(*Initializer),
			}
			yyVAL.node = lhs
			d := lhs.Declarator
			lhs.Initializer.typeCheck(&d.Type, d.Type, lhs.Declarator.specifier.IsStatic(), lx)
			if d.Type.Specifier().IsExtern() {
				id, _ := d.Identifier()
				lx.report.Err(d.Pos(), "'%s' initialized and declared 'extern'", dict.S(id))
			}
		}
	case 95:
		{
			lhs := &StorageClassSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saTypedef
		}
	case 96:
		{
			lhs := &StorageClassSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saExtern
		}
	case 97:
		{
			lhs := &StorageClassSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saStatic
		}
	case 98:
		{
			lhs := &StorageClassSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saAuto
		}
	case 99:
		{
			lhs := &StorageClassSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRegister
		}
	case 100:
		{
			lhs := &TypeSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsVoid)
		}
	case 101:
		{
			lhs := &TypeSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsChar)
		}
	case 102:
		{
			lhs := &TypeSpecifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsShort)
		}
	case 103:
		{
			lhs := &TypeSpecifier{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsInt)
		}
	case 104:
		{
			lhs := &TypeSpecifier{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsLong)
		}
	case 105:
		{
			lhs := &TypeSpecifier{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsFloat)
		}
	case 106:
		{
			lhs := &TypeSpecifier{
				Case:  6,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsDouble)
		}
	case 107:
		{
			lhs := &TypeSpecifier{
				Case:  7,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsSigned)
		}
	case 108:
		{
			lhs := &TypeSpecifier{
				Case:  8,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsUnsigned)
		}
	case 109:
		{
			lhs := &TypeSpecifier{
				Case:  9,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsBool)
		}
	case 110:
		{
			lhs := &TypeSpecifier{
				Case:  10,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsComplex)
		}
	case 111:
		{
			lhs := &TypeSpecifier{
				Case: 11,
				StructOrUnionSpecifier: yyS[yypt-0].node.(*StructOrUnionSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(lhs.StructOrUnionSpecifier.typeSpecifiers())
		}
	case 112:
		{
			lhs := &TypeSpecifier{
				Case:          12,
				EnumSpecifier: yyS[yypt-0].node.(*EnumSpecifier),
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsEnumSpecifier)
		}
	case 113:
		{
			lx := yylex.(*lexer)
			lhs := &TypeSpecifier{
				Case:  13,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsTypedefName)
			_, lhs.scope = lx.scope.Lookup2(NSIdentifiers, lhs.Token.Val)
		}
	case 114:
		{
			lx := yylex.(*lexer)
			lhs := &TypeSpecifier{
				Case:       14,
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsTypeof)
			_, lhs.Type = lhs.Expression.eval(lx)
		}
	case 115:
		{
			lhs := &TypeSpecifier{
				Case:     15,
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				TypeName: yyS[yypt-1].node.(*TypeName),
				Token3:   yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.typeSpecifier = tsEncode(tsTypeof)
			lhs.Type = undefined
			if t := lhs.TypeName.Type; t != nil {
				lhs.Type = t
			}
		}
	case 116:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-1].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
			lx.pushScope(ScopeMembers)
			lx.scope.isUnion = yyS[yypt-2].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
		}
	case 117:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				StructOrUnion: yyS[yypt-5].node.(*StructOrUnion),
				IdentifierOpt: yyS[yypt-4].node.(*IdentifierOpt),
				Token:         yyS[yypt-3].Token,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList).reverse(),
				Token2:                yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			sc := lx.scope
			lhs.scope = sc
			if sc.bitOffset != 0 {
				finishBitField(lhs, lx)
			}

			i := 0
			var bt Type
			var d *Declarator
			for l := lhs.StructDeclarationList; l != nil; l = l.StructDeclarationList {
				for l := l.StructDeclaration.StructDeclaratorList; l != nil; l = l.StructDeclaratorList {
					switch sd := l.StructDeclarator; sd.Case {
					case 0: // Declarator
						d = sd.Declarator
					case 1: // DeclaratorOpt ':' ConstantExpression
						if o := sd.DeclaratorOpt; o != nil {
							x := o.Declarator
							if x.bitOffset == 0 {
								d = x
								bt = lx.scope.bitFieldTypes[i]
								i++
							}
							x.bitFieldType = bt
						}
					}
				}
			}
			lx.scope.bitFieldTypes = nil

			lhs.alignOf = sc.maxAlign
			switch {
			case sc.isUnion:
				lhs.sizeOf = align(sc.maxSize, sc.maxAlign)
			default:
				off := sc.offset
				lhs.sizeOf = align(sc.offset, sc.maxAlign)
				if d != nil {
					d.padding = lhs.sizeOf - off
				}
			}

			lx.popScope(lhs.Token2)
			if o := lhs.IdentifierOpt; o != nil {
				lx.scope.defineStructTag(o.Token, lhs, lx.report)
			}
		}
	case 118:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				Case:          1,
				StructOrUnion: yyS[yypt-1].node.(*StructOrUnion),
				Token:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lx.scope.declareStructTag(lhs.Token, lx.report)
			lhs.scope = lx.scope
		}
	case 119:
		{
			lx := yylex.(*lexer)
			lhs := &StructOrUnionSpecifier{
				Case:          2,
				StructOrUnion: yyS[yypt-3].node.(*StructOrUnion),
				IdentifierOpt: yyS[yypt-2].node.(*IdentifierOpt),
				Token:         yyS[yypt-1].Token,
				Token2:        yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableEmptyStructs {
				lx.report.Err(lhs.Token.Pos(), "empty structs/unions not allowed")
			}
			if o := yyS[yypt-2].node.(*IdentifierOpt); o != nil {
				lx.scope.declareStructTag(o.Token, lx.report)
			}
			lx.scope.isUnion = yyS[yypt-3].node.(*StructOrUnion).Case == 1 // "union"
			lx.scope.prevStructDeclarator = nil
			lhs.alignOf = 1
			lhs.sizeOf = 0
			if o := lhs.IdentifierOpt; o != nil {
				lx.scope.defineStructTag(o.Token, lhs, lx.report)
			}
		}
	case 120:
		{
			yyVAL.node = &StructOrUnion{
				Token: yyS[yypt-0].Token,
			}
		}
	case 121:
		{
			yyVAL.node = &StructOrUnion{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 122:
		{
			yyVAL.node = &StructDeclarationList{
				StructDeclaration: yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 123:
		{
			yyVAL.node = &StructDeclarationList{
				Case: 1,
				StructDeclarationList: yyS[yypt-1].node.(*StructDeclarationList),
				StructDeclaration:     yyS[yypt-0].node.(*StructDeclaration),
			}
		}
	case 124:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclaration{
				SpecifierQualifierList: yyS[yypt-2].node.(*SpecifierQualifierList),
				StructDeclaratorList:   yyS[yypt-1].node.(*StructDeclaratorList).reverse(),
				Token:                  yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			s := lhs.SpecifierQualifierList
			if k := s.kind(); k != Struct && k != Union {
				break
			}

			d := &Declarator{specifier: s}
			dd := &DirectDeclarator{
				Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
				declarator: d,
				idScope:    lx.scope,
				specifier:  s,
			}
			d.DirectDeclarator = dd
			d.setFull(lx)
			for l := lhs.SpecifierQualifierList; l != nil; {
				ts := l.TypeSpecifier
				if ts != nil && ts.Case == 11 && ts.StructOrUnionSpecifier.Case == 0 { // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
					ts.StructOrUnionSpecifier.declarator = d
					break
				}

				if o := l.SpecifierQualifierListOpt; o != nil {
					l = o.SpecifierQualifierList
					continue
				}

				break
			}
		}
	case 125:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclaration{
				Case: 1,
				SpecifierQualifierList: yyS[yypt-1].node.(*SpecifierQualifierList),
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			s := lhs.SpecifierQualifierList
			if !lx.tweaks.enableAnonymousStructFields {
				lx.report.Err(lhs.Token.Pos(), "unnamed fields not allowed")
			} else if k := s.kind(); k != Struct && k != Union {
				lx.report.Err(lhs.Token.Pos(), "only unnamed structs and unions are allowed")
				break
			}

			d := &Declarator{specifier: s}
			dd := &DirectDeclarator{
				Token:      xc.Token{Char: lex.NewChar(lhs.Pos(), 0)},
				declarator: d,
				idScope:    lx.scope,
				specifier:  s,
			}
			d.DirectDeclarator = dd
			d.setFull(lx)

			// we have no struct declarators to parse, so we have to create the case of one implicit declarator
			// because else the size of anonymous members is not included in the struct size!
			dummy := &StructDeclarator{Declarator: d}
			dummy.post(lx)

			for l := lhs.SpecifierQualifierList; l != nil; {
				ts := l.TypeSpecifier
				if ts != nil && ts.Case == 11 && ts.StructOrUnionSpecifier.Case == 0 { // StructOrUnion IdentifierOpt '{' StructDeclarationList '}'
					ts.StructOrUnionSpecifier.declarator = d
					break
				}

				if o := l.SpecifierQualifierListOpt; o != nil {
					l = o.SpecifierQualifierList
					continue
				}

				break
			}
		}
	case 126:
		{
			yyVAL.node = &StructDeclaration{
				Case: 2,
				StaticAssertDeclaration: yyS[yypt-0].node.(*StaticAssertDeclaration),
			}
		}
	case 127:
		{
			lx := yylex.(*lexer)
			lhs := &SpecifierQualifierList{
				TypeSpecifier:             yyS[yypt-1].node.(*TypeSpecifier),
				SpecifierQualifierListOpt: yyS[yypt-0].node.(*SpecifierQualifierListOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeSpecifier
			b := lhs.SpecifierQualifierListOpt
			if b == nil {
				lhs.typeSpecifier = a.typeSpecifier
				break
			}

			lhs.attr = b.attr
			ts := tsEncode(append(tsDecode(a.typeSpecifier), tsDecode(b.typeSpecifier)...)...)
			if _, ok := tsValid[ts]; !ok {
				lx.report.Err(a.Pos(), "invalid type specifier")
				break
			}

			lhs.typeSpecifier = ts
		}
	case 128:
		{
			lx := yylex.(*lexer)
			lhs := &SpecifierQualifierList{
				Case:                      1,
				TypeQualifier:             yyS[yypt-1].node.(*TypeQualifier),
				SpecifierQualifierListOpt: yyS[yypt-0].node.(*SpecifierQualifierListOpt),
			}
			yyVAL.node = lhs
			lx.scope.specifier = lhs
			a := lhs.TypeQualifier
			b := lhs.SpecifierQualifierListOpt
			if b == nil {
				lhs.attr = a.attr
				break
			}

			if a.attr&b.attr != 0 {
				lx.report.Err(a.Pos(), "invalid type qualifier")
				break
			}

			lhs.attr = a.attr | b.attr
			lhs.typeSpecifier = b.typeSpecifier
		}
	case 129:
		{
			yyVAL.node = (*SpecifierQualifierListOpt)(nil)
		}
	case 130:
		{
			lhs := &SpecifierQualifierListOpt{
				SpecifierQualifierList: yyS[yypt-0].node.(*SpecifierQualifierList),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.SpecifierQualifierList.attr
			lhs.typeSpecifier = lhs.SpecifierQualifierList.typeSpecifier
		}
	case 131:
		{
			yyVAL.node = &StructDeclaratorList{
				StructDeclarator: yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 132:
		{
			yyVAL.node = &StructDeclaratorList{
				Case:                 1,
				StructDeclaratorList: yyS[yypt-2].node.(*StructDeclaratorList),
				Token:                yyS[yypt-1].Token,
				StructDeclarator:     yyS[yypt-0].node.(*StructDeclarator),
			}
		}
	case 133:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.post(lx)
		}
	case 134:
		{
			lx := yylex.(*lexer)
			lhs := &StructDeclarator{
				Case:               1,
				DeclaratorOpt:      yyS[yypt-2].node.(*DeclaratorOpt),
				Token:              yyS[yypt-1].Token,
				ConstantExpression: yyS[yypt-0].node.(*ConstantExpression),
			}
			yyVAL.node = lhs
			m := lx.model
			e := lhs.ConstantExpression
			if e.Value == nil {
				e.Value, e.Type = m.value2(1, m.IntType)
			}
			if !IsIntType(e.Type) {
				lx.report.Err(e.Pos(), "bit field width not an integer (have '%s')", e.Type)
				e.Value, e.Type = m.value2(1, m.IntType)
			}
			if o := lhs.DeclaratorOpt; o != nil {
				o.Declarator.setFull(lx)
			}
			lhs.post(lx)
		}
	case 135:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 136:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 137:
		{
			lx := yylex.(*lexer)
			if o := yyS[yypt-0].node.(*IdentifierOpt); o != nil {
				lx.scope.declareEnumTag(o.Token, lx.report)
			}
			lx.iota = 0
		}
	case 138:
		{
			lx := yylex.(*lexer)
			lhs := &EnumSpecifier{
				Token:          yyS[yypt-6].Token,
				IdentifierOpt:  yyS[yypt-5].node.(*IdentifierOpt),
				Token2:         yyS[yypt-3].Token,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList).reverse(),
				CommaOpt:       yyS[yypt-1].node.(*CommaOpt),
				Token3:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if o := lhs.IdentifierOpt; o != nil {
				lx.scope.defineEnumTag(o.Token, lhs, lx.report)
			}
			if !lx.tweaks.enableUnsignedEnums {
				break
			}

			lhs.unsigned = true
		loop:
			for l := lhs.EnumeratorList; l != nil; l = l.EnumeratorList {
				switch e := l.Enumerator; x := e.Value.(type) {
				case int32:
					if x < 0 {
						lhs.unsigned = false
						break loop
					}
				case int64:
					if x < 0 {
						lhs.unsigned = false
						break loop
					}
				default:
					panic(fmt.Errorf("%s: TODO Enumerator.Value type %T", position(e.Pos()), x))
				}
			}
		}
	case 139:
		{
			lx := yylex.(*lexer)
			lhs := &EnumSpecifier{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lx.scope.declareEnumTag(lhs.Token2, lx.report)
		}
	case 140:
		{
			yyVAL.node = &EnumeratorList{
				Enumerator: yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 141:
		{
			yyVAL.node = &EnumeratorList{
				Case:           1,
				EnumeratorList: yyS[yypt-2].node.(*EnumeratorList),
				Token:          yyS[yypt-1].Token,
				Enumerator:     yyS[yypt-0].node.(*Enumerator),
			}
		}
	case 142:
		{
			lx := yylex.(*lexer)
			lhs := &Enumerator{
				EnumerationConstant: yyS[yypt-0].node.(*EnumerationConstant),
			}
			yyVAL.node = lhs
			m := lx.model
			v := m.MustConvert(lx.iota, m.IntType)
			lhs.Value = v
			lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, v)
		}
	case 143:
		{
			lx := yylex.(*lexer)
			lhs := &Enumerator{
				Case:                1,
				EnumerationConstant: yyS[yypt-2].node.(*EnumerationConstant),
				Token:               yyS[yypt-1].Token,
				ConstantExpression:  yyS[yypt-0].node.(*ConstantExpression),
			}
			yyVAL.node = lhs
			m := lx.model
			e := lhs.ConstantExpression
			var v interface{}
			// [0], 6.7.2.2
			// The expression that defines the value of an enumeration
			// constant shall be an integer constant expression that has a
			// value representable as an int.
			switch {
			case !IsIntType(e.Type):
				lx.report.Err(e.Pos(), "not an integer constant expression (have '%s')", e.Type)
				v = m.MustConvert(int32(0), m.IntType)
			default:
				var ok bool
				if v, ok = m.enumValueToInt(e.Value); !ok {
					lx.report.Err(e.Pos(), "overflow in enumeration value: %v", e.Value)
				}
			}

			lhs.Value = v
			lx.scope.defineEnumConst(lx, lhs.EnumerationConstant.Token, v)
		}
	case 144:
		{
			lhs := &TypeQualifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saConst
		}
	case 145:
		{
			lhs := &TypeQualifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saRestrict
		}
	case 146:
		{
			lhs := &TypeQualifier{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saVolatile
		}
	case 147:
		{
			lhs := &FunctionSpecifier{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saInline
		}
	case 148:
		{
			lhs := &FunctionSpecifier{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.attr = saNoreturn
		}
	case 149:
		{
			lx := yylex.(*lexer)
			lhs := &Declarator{
				PointerOpt:       yyS[yypt-1].node.(*PointerOpt),
				DirectDeclarator: yyS[yypt-0].node.(*DirectDeclarator),
			}
			yyVAL.node = lhs
			lhs.specifier = lx.scope.specifier
			lhs.DirectDeclarator.declarator = lhs
		}
	case 150:
		{
			yyVAL.node = (*DeclaratorOpt)(nil)
		}
	case 151:
		{
			yyVAL.node = &DeclaratorOpt{
				Declarator: yyS[yypt-0].node.(*Declarator),
			}
		}
	case 152:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.specifier = lx.scope.specifier
			lx.scope.declareIdentifier(lhs.Token, lhs, lx.report)
			lhs.idScope = lx.scope
		}
	case 153:
		{
			lhs := &DirectDeclarator{
				Case:       1,
				Token:      yyS[yypt-2].Token,
				Declarator: yyS[yypt-1].node.(*Declarator),
				Token2:     yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Declarator.specifier = nil
			lhs.Declarator.DirectDeclarator.parent = lhs
		}
	case 154:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                 2,
				DirectDeclarator:     yyS[yypt-4].node.(*DirectDeclarator),
				Token:                yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				ExpressionOpt:        yyS[yypt-1].node.(*ExpressionOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.elements = -1
			if o := lhs.ExpressionOpt; o != nil {
				var err error
				if lhs.elements, err = elements(o.Expression.eval(lx)); err != nil {
					lx.report.Err(o.Expression.Pos(), "%s", err)
				}

			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 155:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:                 3,
				DirectDeclarator:     yyS[yypt-5].node.(*DirectDeclarator),
				Token:                yyS[yypt-4].Token,
				Token2:               yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Expression:           yyS[yypt-1].node.(*Expression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			var err error
			if lhs.elements, err = elements(lhs.Expression.eval(lx)); err != nil {
				lx.report.Err(lhs.Expression.Pos(), "%s", err)
			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 156:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:              4,
				DirectDeclarator:  yyS[yypt-5].node.(*DirectDeclarator),
				Token:             yyS[yypt-4].Token,
				TypeQualifierList: yyS[yypt-3].node.(*TypeQualifierList).reverse(),
				Token2:            yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token3:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			var err error
			if lhs.elements, err = elements(lhs.Expression.eval(lx)); err != nil {
				lx.report.Err(lhs.Expression.Pos(), "%s", err)
			}
			lhs.DirectDeclarator.parent = lhs
		}
	case 157:
		{
			lhs := &DirectDeclarator{
				Case:                 5,
				DirectDeclarator:     yyS[yypt-4].node.(*DirectDeclarator),
				Token:                yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Token2:               yyS[yypt-1].Token,
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.DirectDeclarator.parent = lhs
			lhs.elements = -1
		}
	case 158:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 159:
		{
			lx := yylex.(*lexer)
			lhs := &DirectDeclarator{
				Case:              6,
				DirectDeclarator:  yyS[yypt-4].node.(*DirectDeclarator),
				Token:             yyS[yypt-3].Token,
				ParameterTypeList: yyS[yypt-1].node.(*ParameterTypeList),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScope(lhs.Token2)
			lhs.DirectDeclarator.parent = lhs
		}
	case 160:
		{
			lhs := &DirectDeclarator{
				Case:              7,
				DirectDeclarator:  yyS[yypt-3].node.(*DirectDeclarator),
				Token:             yyS[yypt-2].Token,
				IdentifierListOpt: yyS[yypt-1].node.(*IdentifierListOpt),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.DirectDeclarator.parent = lhs
		}
	case 161:
		{
			yyVAL.node = &Pointer{
				Token:                yyS[yypt-1].Token,
				TypeQualifierListOpt: yyS[yypt-0].node.(*TypeQualifierListOpt),
			}
		}
	case 162:
		{
			yyVAL.node = &Pointer{
				Case:                 1,
				Token:                yyS[yypt-2].Token,
				TypeQualifierListOpt: yyS[yypt-1].node.(*TypeQualifierListOpt),
				Pointer:              yyS[yypt-0].node.(*Pointer),
			}
		}
	case 163:
		{
			yyVAL.node = (*PointerOpt)(nil)
		}
	case 164:
		{
			yyVAL.node = &PointerOpt{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
		}
	case 165:
		{
			lhs := &TypeQualifierList{
				TypeQualifier: yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			lhs.attr = lhs.TypeQualifier.attr
		}
	case 166:
		{
			lx := yylex.(*lexer)
			lhs := &TypeQualifierList{
				Case:              1,
				TypeQualifierList: yyS[yypt-1].node.(*TypeQualifierList),
				TypeQualifier:     yyS[yypt-0].node.(*TypeQualifier),
			}
			yyVAL.node = lhs
			a := lhs.TypeQualifierList
			b := lhs.TypeQualifier
			if a.attr&b.attr != 0 {
				lx.report.Err(b.Pos(), "invalid type qualifier")
				break
			}

			lhs.attr = a.attr | b.attr
		}
	case 167:
		{
			yyVAL.node = (*TypeQualifierListOpt)(nil)
		}
	case 168:
		{
			yyVAL.node = &TypeQualifierListOpt{
				TypeQualifierList: yyS[yypt-0].node.(*TypeQualifierList).reverse(),
			}
		}
	case 169:
		{
			lhs := &ParameterTypeList{
				ParameterList: yyS[yypt-0].node.(*ParameterList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 170:
		{
			lhs := &ParameterTypeList{
				Case:          1,
				ParameterList: yyS[yypt-2].node.(*ParameterList).reverse(),
				Token:         yyS[yypt-1].Token,
				Token2:        yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.post()
		}
	case 171:
		{
			yyVAL.node = (*ParameterTypeListOpt)(nil)
		}
	case 172:
		{
			yyVAL.node = &ParameterTypeListOpt{
				ParameterTypeList: yyS[yypt-0].node.(*ParameterTypeList),
			}
		}
	case 173:
		{
			yyVAL.node = &ParameterList{
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 174:
		{
			yyVAL.node = &ParameterList{
				Case:                 1,
				ParameterList:        yyS[yypt-2].node.(*ParameterList),
				Token:                yyS[yypt-1].Token,
				ParameterDeclaration: yyS[yypt-0].node.(*ParameterDeclaration),
			}
		}
	case 175:
		{
			lx := yylex.(*lexer)
			lhs := &ParameterDeclaration{
				DeclarationSpecifiers: yyS[yypt-1].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-0].node.(*Declarator),
			}
			yyVAL.node = lhs
			lhs.Declarator.setFull(lx)
			lhs.declarator = lhs.Declarator
		}
	case 176:
		{
			lx := yylex.(*lexer)
			lhs := &ParameterDeclaration{
				Case: 1,
				DeclarationSpecifiers: yyS[yypt-1].node.(*DeclarationSpecifiers),
				AbstractDeclaratorOpt: yyS[yypt-0].node.(*AbstractDeclaratorOpt),
			}
			yyVAL.node = lhs
			if o := lhs.AbstractDeclaratorOpt; o != nil {
				lhs.declarator = o.AbstractDeclarator.declarator
				lhs.declarator.setFull(lx)
				break
			}

			d := &Declarator{
				specifier: lx.scope.specifier,
				DirectDeclarator: &DirectDeclarator{
					Case: 0, // IDENTIFIER
				},
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
			d.setFull(lx)
		}
	case 177:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 178:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 179:
		{
			yyVAL.node = (*IdentifierListOpt)(nil)
		}
	case 180:
		{
			yyVAL.node = &IdentifierListOpt{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 181:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 182:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 183:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeBlock)
		}
	case 184:
		{
			lx := yylex.(*lexer)
			lhs := &TypeName{
				SpecifierQualifierList: yyS[yypt-1].node.(*SpecifierQualifierList),
				AbstractDeclaratorOpt:  yyS[yypt-0].node.(*AbstractDeclaratorOpt),
			}
			yyVAL.node = lhs
			if o := lhs.AbstractDeclaratorOpt; o != nil {
				lhs.declarator = o.AbstractDeclarator.declarator
			} else {
				d := &Declarator{
					specifier: lhs.SpecifierQualifierList,
					DirectDeclarator: &DirectDeclarator{
						Case:    0, // IDENTIFIER
						idScope: lx.scope,
					},
				}
				d.DirectDeclarator.declarator = d
				lhs.declarator = d
			}
			lhs.Type = lhs.declarator.setFull(lx)
			lhs.scope = lx.scope
			lx.popScope(xc.Token{})
		}
	case 185:
		{
			lx := yylex.(*lexer)
			lhs := &AbstractDeclarator{
				Pointer: yyS[yypt-0].node.(*Pointer),
			}
			yyVAL.node = lhs
			d := &Declarator{
				specifier: lx.scope.specifier,
				PointerOpt: &PointerOpt{
					Pointer: lhs.Pointer,
				},
				DirectDeclarator: &DirectDeclarator{
					Case:    0, // IDENTIFIER
					idScope: lx.scope,
				},
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
		}
	case 186:
		{
			lx := yylex.(*lexer)
			lhs := &AbstractDeclarator{
				Case:                     1,
				PointerOpt:               yyS[yypt-1].node.(*PointerOpt),
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
			yyVAL.node = lhs
			d := &Declarator{
				specifier:        lx.scope.specifier,
				PointerOpt:       lhs.PointerOpt,
				DirectDeclarator: lhs.DirectAbstractDeclarator.directDeclarator,
			}
			d.DirectDeclarator.declarator = d
			lhs.declarator = d
		}
	case 187:
		{
			yyVAL.node = (*AbstractDeclaratorOpt)(nil)
		}
	case 188:
		{
			yyVAL.node = &AbstractDeclaratorOpt{
				AbstractDeclarator: yyS[yypt-0].node.(*AbstractDeclarator),
			}
		}
	case 189:
		{
			lhs := &DirectAbstractDeclarator{
				Token:              yyS[yypt-2].Token,
				AbstractDeclarator: yyS[yypt-1].node.(*AbstractDeclarator),
				Token2:             yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.AbstractDeclarator.declarator.specifier = nil
			lhs.directDeclarator = &DirectDeclarator{
				Case:       1, // '(' Declarator ')'
				Declarator: lhs.AbstractDeclarator.declarator,
			}
			lhs.AbstractDeclarator.declarator.DirectDeclarator.parent = lhs.directDeclarator
		}
	case 190:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 1,
				DirectAbstractDeclaratorOpt: yyS[yypt-3].node.(*DirectAbstractDeclaratorOpt),
				Token:         yyS[yypt-2].Token,
				ExpressionOpt: yyS[yypt-1].node.(*ExpressionOpt),
				Token2:        yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			nElements := -1
			if o := lhs.ExpressionOpt; o != nil {
				var err error
				if nElements, err = elements(o.Expression.eval(lx)); err != nil {
					lx.report.Err(o.Expression.Pos(), "%s", err)
				}
			}
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:             2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
				DirectDeclarator: dd,
				ExpressionOpt:    lhs.ExpressionOpt,
				elements:         nElements,
			}
			dd.parent = lhs.directDeclarator
		}
	case 191:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 2,
				DirectAbstractDeclaratorOpt: yyS[yypt-4].node.(*DirectAbstractDeclaratorOpt),
				Token:             yyS[yypt-3].Token,
				TypeQualifierList: yyS[yypt-2].node.(*TypeQualifierList).reverse(),
				ExpressionOpt:     yyS[yypt-1].node.(*ExpressionOpt),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if o := lhs.ExpressionOpt; o != nil {
				o.Expression.eval(lx)
			}
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:                 2, // DirectDeclarator '[' TypeQualifierListOpt ExpressionOpt ']'
				DirectDeclarator:     dd,
				TypeQualifierListOpt: &TypeQualifierListOpt{lhs.TypeQualifierList},
				ExpressionOpt:        lhs.ExpressionOpt,
			}
			dd.parent = lhs.directDeclarator
		}
	case 192:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 3,
				DirectAbstractDeclaratorOpt: yyS[yypt-5].node.(*DirectAbstractDeclaratorOpt),
				Token:                yyS[yypt-4].Token,
				Token2:               yyS[yypt-3].Token,
				TypeQualifierListOpt: yyS[yypt-2].node.(*TypeQualifierListOpt),
				Expression:           yyS[yypt-1].node.(*Expression),
				Token3:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:                 2, // DirectDeclarator '[' "static" TypeQualifierListOpt Expression ']'
				DirectDeclarator:     dd,
				TypeQualifierListOpt: lhs.TypeQualifierListOpt,
				Expression:           lhs.Expression,
			}
			dd.parent = lhs.directDeclarator
		}
	case 193:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 4,
				DirectAbstractDeclaratorOpt: yyS[yypt-5].node.(*DirectAbstractDeclaratorOpt),
				Token:             yyS[yypt-4].Token,
				TypeQualifierList: yyS[yypt-3].node.(*TypeQualifierList).reverse(),
				Token2:            yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token3:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:              4, // DirectDeclarator '[' TypeQualifierList "static" Expression ']'
				DirectDeclarator:  dd,
				TypeQualifierList: lhs.TypeQualifierList,
				Expression:        lhs.Expression,
			}
			dd.parent = lhs.directDeclarator
		}
	case 194:
		{
			lhs := &DirectAbstractDeclarator{
				Case: 5,
				DirectAbstractDeclaratorOpt: yyS[yypt-3].node.(*DirectAbstractDeclaratorOpt),
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			var dd *DirectDeclarator
			switch o := lhs.DirectAbstractDeclaratorOpt; {
			case o == nil:
				dd = &DirectDeclarator{
					Case: 0, // IDENTIFIER
				}
			default:
				dd = o.DirectAbstractDeclarator.directDeclarator
			}
			lhs.directDeclarator = &DirectDeclarator{
				Case:             5, // DirectDeclarator '[' TypeQualifierListOpt '*' ']'
				DirectDeclarator: dd,
			}
			dd.parent = lhs.directDeclarator
		}
	case 195:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 196:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case:                 6,
				Token:                yyS[yypt-3].Token,
				ParameterTypeListOpt: yyS[yypt-1].node.(*ParameterTypeListOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScope(lhs.Token2)
			switch o := lhs.ParameterTypeListOpt; {
			case o != nil:
				lhs.directDeclarator = &DirectDeclarator{
					Case: 6, // DirectDeclarator '(' ParameterTypeList ')'
					DirectDeclarator: &DirectDeclarator{
						Case: 0, // IDENTIFIER
					},
					ParameterTypeList: o.ParameterTypeList,
				}
			default:
				lhs.directDeclarator = &DirectDeclarator{
					Case: 7, // DirectDeclarator '(' IdentifierListOpt ')'
					DirectDeclarator: &DirectDeclarator{
						Case: 0, // IDENTIFIER
					},
				}
			}
			lhs.directDeclarator.DirectDeclarator.parent = lhs.directDeclarator
		}
	case 197:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 198:
		{
			lx := yylex.(*lexer)
			lhs := &DirectAbstractDeclarator{
				Case: 7,
				DirectAbstractDeclarator: yyS[yypt-4].node.(*DirectAbstractDeclarator),
				Token:                yyS[yypt-3].Token,
				ParameterTypeListOpt: yyS[yypt-1].node.(*ParameterTypeListOpt),
				Token2:               yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScope(lhs.Token2)
			switch o := lhs.ParameterTypeListOpt; {
			case o != nil:
				lhs.directDeclarator = &DirectDeclarator{
					Case:              6, // DirectDeclarator '(' ParameterTypeList ')'
					DirectDeclarator:  lhs.DirectAbstractDeclarator.directDeclarator,
					ParameterTypeList: o.ParameterTypeList,
				}
			default:
				lhs.directDeclarator = &DirectDeclarator{
					Case:             7, // DirectDeclarator '(' IdentifierListOpt ')'
					DirectDeclarator: lhs.DirectAbstractDeclarator.directDeclarator,
				}
			}
			lhs.directDeclarator.DirectDeclarator.parent = lhs.directDeclarator
		}
	case 199:
		{
			yyVAL.node = (*DirectAbstractDeclaratorOpt)(nil)
		}
	case 200:
		{
			yyVAL.node = &DirectAbstractDeclaratorOpt{
				DirectAbstractDeclarator: yyS[yypt-0].node.(*DirectAbstractDeclarator),
			}
		}
	case 201:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lhs.Expression.eval(lx)
		}
	case 202:
		{
			yyVAL.node = &Initializer{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				InitializerList: yyS[yypt-2].node.(*InitializerList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 203:
		{
			lx := yylex.(*lexer)
			lhs := &Initializer{
				Case:        2,
				Token:       yyS[yypt-2].Token,
				Token2:      yyS[yypt-1].Token,
				Initializer: yyS[yypt-0].node.(*Initializer),
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableLegacyDesignators {
				lx.report.Err(lhs.Pos(), "legacy designators not enabled")
			}
		}
	case 204:
		{
			yyVAL.node = &InitializerList{
				DesignationOpt: yyS[yypt-1].node.(*DesignationOpt),
				Initializer:    yyS[yypt-0].node.(*Initializer),
			}
		}
	case 205:
		{
			yyVAL.node = &InitializerList{
				Case:            1,
				InitializerList: yyS[yypt-3].node.(*InitializerList),
				Token:           yyS[yypt-2].Token,
				DesignationOpt:  yyS[yypt-1].node.(*DesignationOpt),
				Initializer:     yyS[yypt-0].node.(*Initializer),
			}
		}
	case 206:
		{
			yyVAL.node = (*InitializerList)(nil)
		}
	case 207:
		{
			yyVAL.node = &Designation{
				DesignatorList: yyS[yypt-1].node.(*DesignatorList).reverse(),
				Token:          yyS[yypt-0].Token,
			}
		}
	case 208:
		{
			yyVAL.node = (*DesignationOpt)(nil)
		}
	case 209:
		{
			yyVAL.node = &DesignationOpt{
				Designation: yyS[yypt-0].node.(*Designation),
			}
		}
	case 210:
		{
			yyVAL.node = &DesignatorList{
				Designator: yyS[yypt-0].node.(*Designator),
			}
		}
	case 211:
		{
			yyVAL.node = &DesignatorList{
				Case:           1,
				DesignatorList: yyS[yypt-1].node.(*DesignatorList),
				Designator:     yyS[yypt-0].node.(*Designator),
			}
		}
	case 212:
		{
			yyVAL.node = &Designator{
				Token:              yyS[yypt-2].Token,
				ConstantExpression: yyS[yypt-1].node.(*ConstantExpression),
				Token2:             yyS[yypt-0].Token,
			}
		}
	case 213:
		{
			yyVAL.node = &Designator{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 214:
		{
			yyVAL.node = &Statement{
				LabeledStatement: yyS[yypt-0].node.(*LabeledStatement),
			}
		}
	case 215:
		{
			yyVAL.node = &Statement{
				Case:              1,
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
		}
	case 216:
		{
			yyVAL.node = &Statement{
				Case:                2,
				ExpressionStatement: yyS[yypt-0].node.(*ExpressionStatement),
			}
		}
	case 217:
		{
			yyVAL.node = &Statement{
				Case:               3,
				SelectionStatement: yyS[yypt-0].node.(*SelectionStatement),
			}
		}
	case 218:
		{
			yyVAL.node = &Statement{
				Case:               4,
				IterationStatement: yyS[yypt-0].node.(*IterationStatement),
			}
		}
	case 219:
		{
			yyVAL.node = &Statement{
				Case:          5,
				JumpStatement: yyS[yypt-0].node.(*JumpStatement),
			}
		}
	case 220:
		{
			yyVAL.node = &Statement{
				Case:               6,
				AssemblerStatement: yyS[yypt-0].node.(*AssemblerStatement),
			}
		}
	case 221:
		{
			yyVAL.node = &LabeledStatement{
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 222:
		{
			yyVAL.node = &LabeledStatement{
				Case:               1,
				Token:              yyS[yypt-3].Token,
				ConstantExpression: yyS[yypt-2].node.(*ConstantExpression),
				Token2:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 223:
		{
			yyVAL.node = &LabeledStatement{
				Case:      2,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 224:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 225:
		{
			lx := yylex.(*lexer)
			lhs := &CompoundStatement{
				Token:            yyS[yypt-3].Token,
				BlockItemListOpt: yyS[yypt-1].node.(*BlockItemListOpt),
				Token2:           yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.scope = lx.scope
			lx.popScope(lhs.Token2)
		}
	case 226:
		{
			yyVAL.node = &BlockItemList{
				BlockItem: yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 227:
		{
			yyVAL.node = &BlockItemList{
				Case:          1,
				BlockItemList: yyS[yypt-1].node.(*BlockItemList),
				BlockItem:     yyS[yypt-0].node.(*BlockItem),
			}
		}
	case 228:
		{
			yyVAL.node = (*BlockItemListOpt)(nil)
		}
	case 229:
		{
			yyVAL.node = &BlockItemListOpt{
				BlockItemList: yyS[yypt-0].node.(*BlockItemList).reverse(),
			}
		}
	case 230:
		{
			yyVAL.node = &BlockItem{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 231:
		{
			yyVAL.node = &BlockItem{
				Case:      1,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 232:
		{
			yyVAL.node = &ExpressionStatement{
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token:             yyS[yypt-0].Token,
			}
		}
	case 233:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 234:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Case:           1,
				Token:          yyS[yypt-6].Token,
				Token2:         yyS[yypt-5].Token,
				ExpressionList: yyS[yypt-4].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-3].Token,
				Statement:      yyS[yypt-2].node.(*Statement),
				Token4:         yyS[yypt-1].Token,
				Statement2:     yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 235:
		{
			lx := yylex.(*lexer)
			lhs := &SelectionStatement{
				Case:           2,
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 236:
		{
			lx := yylex.(*lexer)
			lhs := &IterationStatement{
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Statement:      yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 237:
		{
			lx := yylex.(*lexer)
			lhs := &IterationStatement{
				Case:           1,
				Token:          yyS[yypt-6].Token,
				Statement:      yyS[yypt-5].node.(*Statement),
				Token2:         yyS[yypt-4].Token,
				Token3:         yyS[yypt-3].Token,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token4:         yyS[yypt-1].Token,
				Token5:         yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.ExpressionList.eval(lx)
		}
	case 238:
		{
			yyVAL.node = &IterationStatement{
				Case:               2,
				Token:              yyS[yypt-8].Token,
				Token2:             yyS[yypt-7].Token,
				ExpressionListOpt:  yyS[yypt-6].node.(*ExpressionListOpt),
				Token3:             yyS[yypt-5].Token,
				ExpressionListOpt2: yyS[yypt-4].node.(*ExpressionListOpt),
				Token4:             yyS[yypt-3].Token,
				ExpressionListOpt3: yyS[yypt-2].node.(*ExpressionListOpt),
				Token5:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 239:
		{
			yyVAL.node = &IterationStatement{
				Case:               3,
				Token:              yyS[yypt-7].Token,
				Token2:             yyS[yypt-6].Token,
				Declaration:        yyS[yypt-5].node.(*Declaration),
				ExpressionListOpt:  yyS[yypt-4].node.(*ExpressionListOpt),
				Token3:             yyS[yypt-3].Token,
				ExpressionListOpt2: yyS[yypt-2].node.(*ExpressionListOpt),
				Token4:             yyS[yypt-1].Token,
				Statement:          yyS[yypt-0].node.(*Statement),
			}
		}
	case 240:
		{
			yyVAL.node = &JumpStatement{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 241:
		{
			yyVAL.node = &JumpStatement{
				Case:   1,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 242:
		{
			yyVAL.node = &JumpStatement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 243:
		{
			yyVAL.node = &JumpStatement{
				Case:              3,
				Token:             yyS[yypt-2].Token,
				ExpressionListOpt: yyS[yypt-1].node.(*ExpressionListOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 244:
		{
			lx := yylex.(*lexer)
			lhs := &JumpStatement{
				Case:       4,
				Token:      yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token2:     yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			_, t := lhs.Expression.eval(lx)
			if t == nil {
				break
			}

			for t != nil && t.Kind() == Ptr {
				t = t.Element()
			}

			if t == nil || t.Kind() != Void {
				lx.report.Err(lhs.Pos(), "invalid computed goto argument type, have '%s'", t)
			}

			if !lx.tweaks.enableComputedGotos {
				lx.report.Err(lhs.Pos(), "computed gotos not enabled")
			}
		}
	case 245:
		{
			yyVAL.node = &TranslationUnit{
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 246:
		{
			yyVAL.node = &TranslationUnit{
				Case:                1,
				TranslationUnit:     yyS[yypt-1].node.(*TranslationUnit),
				ExternalDeclaration: yyS[yypt-0].node.(*ExternalDeclaration),
			}
		}
	case 247:
		{
			yyVAL.node = &ExternalDeclaration{
				FunctionDefinition: yyS[yypt-0].node.(*FunctionDefinition),
			}
		}
	case 248:
		{
			yyVAL.node = &ExternalDeclaration{
				Case:        1,
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 249:
		{
			yyVAL.node = &ExternalDeclaration{
				Case: 2,
				BasicAssemblerStatement: yyS[yypt-1].node.(*BasicAssemblerStatement),
				Token: yyS[yypt-0].Token,
			}
		}
	case 250:
		{
			lx := yylex.(*lexer)
			lhs := &ExternalDeclaration{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableEmptyDeclarations {
				lx.report.Err(lhs.Pos(), "C++11 empty declarations are illegal in C99.")
			}
		}
	case 251:
		{
			lx := yylex.(*lexer)
			if ds := yyS[yypt-2].node.(*DeclarationSpecifiers); ds.typeSpecifier == 0 {
				ds.typeSpecifier = tsEncode(tsInt)
				yyS[yypt-1].node.(*Declarator).Type = lx.model.IntType
				if !lx.tweaks.enableOmitFuncRetType {
					lx.report.Err(yyS[yypt-1].node.Pos(), "missing function return type")
				}
			}
			var fd *FunctionDefinition
			fd.post(lx, yyS[yypt-1].node.(*Declarator), yyS[yypt-0].node.(*DeclarationListOpt))
		}
	case 252:
		{
			yyVAL.node = &FunctionDefinition{
				DeclarationSpecifiers: yyS[yypt-4].node.(*DeclarationSpecifiers),
				Declarator:            yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt:    yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:          yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 253:
		{
			lx := yylex.(*lexer)
			lx.scope.specifier = &DeclarationSpecifiers{typeSpecifier: tsEncode(tsInt)}
		}
	case 254:
		{
			lx := yylex.(*lexer)
			if !lx.tweaks.enableOmitFuncRetType {
				lx.report.Err(yyS[yypt-1].node.Pos(), "missing function return type")
			}
			var fd *FunctionDefinition
			fd.post(lx, yyS[yypt-1].node.(*Declarator), yyS[yypt-0].node.(*DeclarationListOpt))
		}
	case 255:
		{
			yyVAL.node = &FunctionDefinition{
				Case:               1,
				Declarator:         yyS[yypt-3].node.(*Declarator),
				DeclarationListOpt: yyS[yypt-2].node.(*DeclarationListOpt),
				FunctionBody:       yyS[yypt-0].node.(*FunctionBody),
			}
		}
	case 256:
		{
			lx := yylex.(*lexer)
			// Handle __func__, [0], 6.4.2.2.
			id, _ := lx.fnDeclarator.Identifier()
			lx.injectFunc = []xc.Token{
				{lex.Char{Rune: STATIC}, idStatic},
				{lex.Char{Rune: CONST}, idConst},
				{lex.Char{Rune: CHAR}, idChar},
				{lex.Char{Rune: IDENTIFIER}, idMagicFunc},
				{lex.Char{Rune: '['}, 0},
				{lex.Char{Rune: ']'}, 0},
				{lex.Char{Rune: '='}, 0},
				{lex.Char{Rune: STRINGLITERAL}, xc.Dict.SID(fmt.Sprintf("%q", xc.Dict.S(id)))},
				{lex.Char{Rune: ';'}, 0},
			}
		}
	case 257:
		{
			lhs := &FunctionBody{
				CompoundStatement: yyS[yypt-0].node.(*CompoundStatement),
			}
			yyVAL.node = lhs
			lhs.scope = lhs.CompoundStatement.scope
		}
	case 258:
		{
			lx := yylex.(*lexer)
			m := lx.scope.mergeScope
			lx.pushScope(ScopeBlock)
			if m != nil {
				lx.scope.merge(m)
			}
			lx.scope.mergeScope = nil
		}
	case 259:
		{
			lx := yylex.(*lexer)
			lhs := &FunctionBody{
				Case:               1,
				AssemblerStatement: yyS[yypt-1].node.(*AssemblerStatement),
				Token:              yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.scope = lx.scope
			lx.popScope(lx.tokPrev)
		}
	case 260:
		{
			yyVAL.node = &DeclarationList{
				Declaration: yyS[yypt-0].node.(*Declaration),
			}
		}
	case 261:
		{
			yyVAL.node = &DeclarationList{
				Case:            1,
				DeclarationList: yyS[yypt-1].node.(*DeclarationList),
				Declaration:     yyS[yypt-0].node.(*Declaration),
			}
		}
	case 262:
		{
			yyVAL.node = (*DeclarationListOpt)(nil)
		}
	case 263:
		{
			lx := yylex.(*lexer)
			lx.pushScope(ScopeParams)
		}
	case 264:
		{
			lx := yylex.(*lexer)
			lhs := &DeclarationListOpt{
				DeclarationList: yyS[yypt-0].node.(*DeclarationList).reverse(),
			}
			yyVAL.node = lhs
			lhs.paramsScope, _ = lx.popScopePos(lhs.Pos())
		}
	case 265:
		{
			yyVAL.node = &AssemblerInstructions{
				Token: yyS[yypt-0].Token,
			}
		}
	case 266:
		{
			yyVAL.node = &AssemblerInstructions{
				Case: 1,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions),
				Token: yyS[yypt-0].Token,
			}
		}
	case 267:
		{
			yyVAL.node = &BasicAssemblerStatement{
				Token:                 yyS[yypt-4].Token,
				VolatileOpt:           yyS[yypt-3].node.(*VolatileOpt),
				Token2:                yyS[yypt-2].Token,
				AssemblerInstructions: yyS[yypt-1].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-0].Token,
			}
		}
	case 268:
		{
			yyVAL.node = (*VolatileOpt)(nil)
		}
	case 269:
		{
			yyVAL.node = &VolatileOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 270:
		{
			yyVAL.node = &AssemblerOperand{
				AssemblerSymbolicNameOpt: yyS[yypt-4].node.(*AssemblerSymbolicNameOpt),
				Token:      yyS[yypt-3].Token,
				Token2:     yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token3:     yyS[yypt-0].Token,
			}
		}
	case 271:
		{
			yyVAL.node = &AssemblerOperands{
				AssemblerOperand: yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 272:
		{
			yyVAL.node = &AssemblerOperands{
				Case:              1,
				AssemblerOperands: yyS[yypt-2].node.(*AssemblerOperands),
				Token:             yyS[yypt-1].Token,
				AssemblerOperand:  yyS[yypt-0].node.(*AssemblerOperand),
			}
		}
	case 273:
		{
			yyVAL.node = (*AssemblerSymbolicNameOpt)(nil)
		}
	case 274:
		{
			yyVAL.node = &AssemblerSymbolicNameOpt{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 275:
		{
			yyVAL.node = &Clobbers{
				Token: yyS[yypt-0].Token,
			}
		}
	case 276:
		{
			yyVAL.node = &Clobbers{
				Case:     1,
				Clobbers: yyS[yypt-2].node.(*Clobbers),
				Token:    yyS[yypt-1].Token,
				Token2:   yyS[yypt-0].Token,
			}
		}
	case 277:
		{
			yyVAL.node = &AssemblerStatement{
				BasicAssemblerStatement: yyS[yypt-0].node.(*BasicAssemblerStatement),
			}
		}
	case 278:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  1,
				Token:                 yyS[yypt-6].Token,
				VolatileOpt:           yyS[yypt-5].node.(*VolatileOpt),
				Token2:                yyS[yypt-4].Token,
				AssemblerInstructions: yyS[yypt-3].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-2].Token,
				AssemblerOperands:     yyS[yypt-1].node.(*AssemblerOperands).reverse(),
				Token4:                yyS[yypt-0].Token,
			}
		}
	case 279:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  2,
				Token:                 yyS[yypt-8].Token,
				VolatileOpt:           yyS[yypt-7].node.(*VolatileOpt),
				Token2:                yyS[yypt-6].Token,
				AssemblerInstructions: yyS[yypt-5].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-4].Token,
				AssemblerOperands:     yyS[yypt-3].node.(*AssemblerOperands).reverse(),
				Token4:                yyS[yypt-2].Token,
				AssemblerOperands2:    yyS[yypt-1].node.(*AssemblerOperands).reverse(),
				Token5:                yyS[yypt-0].Token,
			}
		}
	case 280:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  3,
				Token:                 yyS[yypt-10].Token,
				VolatileOpt:           yyS[yypt-9].node.(*VolatileOpt),
				Token2:                yyS[yypt-8].Token,
				AssemblerInstructions: yyS[yypt-7].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-6].Token,
				AssemblerOperands:     yyS[yypt-5].node.(*AssemblerOperands).reverse(),
				Token4:                yyS[yypt-4].Token,
				AssemblerOperands2:    yyS[yypt-3].node.(*AssemblerOperands).reverse(),
				Token5:                yyS[yypt-2].Token,
				Clobbers:              yyS[yypt-1].node.(*Clobbers).reverse(),
				Token6:                yyS[yypt-0].Token,
			}
		}
	case 281:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  4,
				Token:                 yyS[yypt-12].Token,
				VolatileOpt:           yyS[yypt-11].node.(*VolatileOpt),
				Token2:                yyS[yypt-10].Token,
				Token3:                yyS[yypt-9].Token,
				AssemblerInstructions: yyS[yypt-8].node.(*AssemblerInstructions).reverse(),
				Token4:                yyS[yypt-7].Token,
				Token5:                yyS[yypt-6].Token,
				AssemblerOperands:     yyS[yypt-5].node.(*AssemblerOperands).reverse(),
				Token6:                yyS[yypt-4].Token,
				Clobbers:              yyS[yypt-3].node.(*Clobbers).reverse(),
				Token7:                yyS[yypt-2].Token,
				IdentifierList:        yyS[yypt-1].node.(*IdentifierList).reverse(),
				Token8:                yyS[yypt-0].Token,
			}
		}
	case 282:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  5,
				Token:                 yyS[yypt-5].Token,
				VolatileOpt:           yyS[yypt-4].node.(*VolatileOpt),
				Token2:                yyS[yypt-3].Token,
				AssemblerInstructions: yyS[yypt-2].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-1].Token,
				Token4:                yyS[yypt-0].Token,
			}
		}
	case 283:
		{
			yyVAL.node = &AssemblerStatement{
				Case:                  6,
				Token:                 yyS[yypt-7].Token,
				VolatileOpt:           yyS[yypt-6].node.(*VolatileOpt),
				Token2:                yyS[yypt-5].Token,
				AssemblerInstructions: yyS[yypt-4].node.(*AssemblerInstructions).reverse(),
				Token3:                yyS[yypt-3].Token,
				Token4:                yyS[yypt-2].Token,
				AssemblerOperands:     yyS[yypt-1].node.(*AssemblerOperands).reverse(),
				Token5:                yyS[yypt-0].Token,
			}
		}
	case 284:
		{
			lx := yylex.(*lexer)
			lhs := &StaticAssertDeclaration{
				Token:              yyS[yypt-6].Token,
				Token2:             yyS[yypt-5].Token,
				ConstantExpression: yyS[yypt-4].node.(*ConstantExpression),
				Token3:             yyS[yypt-3].Token,
				Token4:             yyS[yypt-2].Token,
				Token5:             yyS[yypt-1].Token,
				Token6:             yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			ce := lhs.ConstantExpression
			if ce.Type == nil || ce.Type.Kind() == Undefined || ce.Value == nil || !IsIntType(ce.Type) {
				lx.report.Err(ce.Pos(), "invalid static assert expression (have '%v')", ce.Type)
				break
			}

			if !isNonZero(ce.Value) {
				lx.report.ErrTok(lhs.Token, "%s", lhs.Token4.S())
			}
		}
	case 285:
		{
			lx := yylex.(*lexer)
			lhs := &PreprocessingFile{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
			yyVAL.node = lhs
			lhs.path = lx.file.Name()
		}
	case 286:
		{
			yyVAL.node = &GroupList{
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 287:
		{
			yyVAL.node = &GroupList{
				Case:      1,
				GroupList: yyS[yypt-1].node.(*GroupList),
				GroupPart: yyS[yypt-0].groupPart,
			}
		}
	case 288:
		{
			yyVAL.node = (*GroupListOpt)(nil)
		}
	case 289:
		{
			yyVAL.node = &GroupListOpt{
				GroupList: yyS[yypt-0].node.(*GroupList).reverse(),
			}
		}
	case 290:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 291:
		{
			yyVAL.groupPart = yyS[yypt-0].node.(Node)
		}
	case 292:
		{
			yyVAL.groupPart = yyS[yypt-2].Token
		}
	case 293:
		{
			yyVAL.groupPart = yyS[yypt-0].toks
		}
	case 294:
		{
			yyVAL.node = &IfSection{
				IfGroup:          yyS[yypt-3].node.(*IfGroup),
				ElifGroupListOpt: yyS[yypt-2].node.(*ElifGroupListOpt),
				ElseGroupOpt:     yyS[yypt-1].node.(*ElseGroupOpt),
				EndifLine:        yyS[yypt-0].node.(*EndifLine),
			}
		}
	case 295:
		{
			yyVAL.node = &IfGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 296:
		{
			yyVAL.node = &IfGroup{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 297:
		{
			yyVAL.node = &IfGroup{
				Case:         2,
				Token:        yyS[yypt-3].Token,
				Token2:       yyS[yypt-2].Token,
				Token3:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 298:
		{
			yyVAL.node = &ElifGroupList{
				ElifGroup: yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 299:
		{
			yyVAL.node = &ElifGroupList{
				Case:          1,
				ElifGroupList: yyS[yypt-1].node.(*ElifGroupList),
				ElifGroup:     yyS[yypt-0].node.(*ElifGroup),
			}
		}
	case 300:
		{
			yyVAL.node = (*ElifGroupListOpt)(nil)
		}
	case 301:
		{
			yyVAL.node = &ElifGroupListOpt{
				ElifGroupList: yyS[yypt-0].node.(*ElifGroupList).reverse(),
			}
		}
	case 302:
		{
			yyVAL.node = &ElifGroup{
				Token:        yyS[yypt-3].Token,
				PPTokenList:  yyS[yypt-2].toks,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 303:
		{
			yyVAL.node = &ElseGroup{
				Token:        yyS[yypt-2].Token,
				Token2:       yyS[yypt-1].Token,
				GroupListOpt: yyS[yypt-0].node.(*GroupListOpt),
			}
		}
	case 304:
		{
			yyVAL.node = (*ElseGroupOpt)(nil)
		}
	case 305:
		{
			yyVAL.node = &ElseGroupOpt{
				ElseGroup: yyS[yypt-0].node.(*ElseGroup),
			}
		}
	case 306:
		{
			yyVAL.node = &EndifLine{
				Token: yyS[yypt-0].Token,
			}
		}
	case 307:
		{
			yyVAL.node = &ControlLine{
				Token:           yyS[yypt-2].Token,
				Token2:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 308:
		{
			yyVAL.node = &ControlLine{
				Case:            1,
				Token:           yyS[yypt-4].Token,
				Token2:          yyS[yypt-3].Token,
				Token3:          yyS[yypt-2].Token,
				Token4:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 309:
		{
			yyVAL.node = &ControlLine{
				Case:            2,
				Token:           yyS[yypt-6].Token,
				Token2:          yyS[yypt-5].Token,
				IdentifierList:  yyS[yypt-4].node.(*IdentifierList).reverse(),
				Token3:          yyS[yypt-3].Token,
				Token4:          yyS[yypt-2].Token,
				Token5:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
		}
	case 310:
		{
			yyVAL.node = &ControlLine{
				Case:              3,
				Token:             yyS[yypt-4].Token,
				Token2:            yyS[yypt-3].Token,
				IdentifierListOpt: yyS[yypt-2].node.(*IdentifierListOpt),
				Token3:            yyS[yypt-1].Token,
				ReplacementList:   yyS[yypt-0].toks,
			}
		}
	case 311:
		{
			yyVAL.node = &ControlLine{
				Case:           4,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 312:
		{
			yyVAL.node = &ControlLine{
				Case:  5,
				Token: yyS[yypt-0].Token,
			}
		}
	case 313:
		{
			yyVAL.node = &ControlLine{
				Case:        6,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 314:
		{
			yyVAL.node = &ControlLine{
				Case:        7,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 315:
		{
			yyVAL.node = &ControlLine{
				Case:           8,
				Token:          yyS[yypt-1].Token,
				PPTokenListOpt: yyS[yypt-0].toks,
			}
		}
	case 316:
		{
			yyVAL.node = &ControlLine{
				Case:   9,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 317:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:            10,
				Token:           yyS[yypt-5].Token,
				Token2:          yyS[yypt-4].Token,
				IdentifierList:  yyS[yypt-3].node.(*IdentifierList).reverse(),
				Token3:          yyS[yypt-2].Token,
				Token4:          yyS[yypt-1].Token,
				ReplacementList: yyS[yypt-0].toks,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableDefineOmitCommaBeforeDDD {
				lx.report.ErrTok(lhs.Token4, "missing comma before \"...\"")
			}
		}
	case 318:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:   11,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if !lx.tweaks.enableEmptyDefine {
				lx.report.ErrTok(lhs.Token2, "expected identifier")
			}
		}
	case 319:
		{
			lx := yylex.(*lexer)
			lhs := &ControlLine{
				Case:        12,
				Token:       yyS[yypt-3].Token,
				Token2:      yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token3:      yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			toks := decodeTokens(lhs.PPTokenList, nil, false)
			if len(toks) == 0 {
				lhs.Case = 9 // PPUNDEF IDENTIFIER '\n'
				break
			}

			lx.report.ErrTok(toks[0], "extra tokens after #undef argument")
		}
	case 320:
		{
			yyVAL.node = &ControlLine{
				Case:        13,
				Token:       yyS[yypt-2].Token,
				PPTokenList: yyS[yypt-1].toks,
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 323:
		{
			lx := yylex.(*lexer)
			yyVAL.toks = PPTokenList(dict.ID(lx.encBuf))
			lx.encBuf = lx.encBuf[:0]
			lx.encPos = 0
		}
	case 324:
		{
			yyVAL.toks = 0
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
