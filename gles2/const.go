// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 04 Jul 2017 10:30:15 CEST.
// By https://git.io/c-for-go. DO NOT EDIT.

package gles2

/*
#cgo LDFLAGS: -lGLESv2
#include <GLES2/gl2.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// ES_VERSION_2_0 as defined in GLES2/gl2.h:54
	ES_VERSION_2_0 = 1
	// DEPTH_BUFFER_BIT as defined in GLES2/gl2.h:76
	DEPTH_BUFFER_BIT = 0x00000100
	// STENCIL_BUFFER_BIT as defined in GLES2/gl2.h:77
	STENCIL_BUFFER_BIT = 0x00000400
	// COLOR_BUFFER_BIT as defined in GLES2/gl2.h:78
	COLOR_BUFFER_BIT = 0x00004000
	// FALSE as defined in GLES2/gl2.h:79
	FALSE = 0
	// TRUE as defined in GLES2/gl2.h:80
	TRUE = 1
	// POINTS as defined in GLES2/gl2.h:81
	POINTS = 0x0000
	// LINES as defined in GLES2/gl2.h:82
	LINES = 0x0001
	// LINE_LOOP as defined in GLES2/gl2.h:83
	LINE_LOOP = 0x0002
	// LINE_STRIP as defined in GLES2/gl2.h:84
	LINE_STRIP = 0x0003
	// TRIANGLES as defined in GLES2/gl2.h:85
	TRIANGLES = 0x0004
	// TRIANGLE_STRIP as defined in GLES2/gl2.h:86
	TRIANGLE_STRIP = 0x0005
	// TRIANGLE_FAN as defined in GLES2/gl2.h:87
	TRIANGLE_FAN = 0x0006
	// ZERO as defined in GLES2/gl2.h:88
	ZERO = 0
	// ONE as defined in GLES2/gl2.h:89
	ONE = 1
	// SRC_COLOR as defined in GLES2/gl2.h:90
	SRC_COLOR = 0x0300
	// ONE_MINUS_SRC_COLOR as defined in GLES2/gl2.h:91
	ONE_MINUS_SRC_COLOR = 0x0301
	// SRC_ALPHA as defined in GLES2/gl2.h:92
	SRC_ALPHA = 0x0302
	// ONE_MINUS_SRC_ALPHA as defined in GLES2/gl2.h:93
	ONE_MINUS_SRC_ALPHA = 0x0303
	// DST_ALPHA as defined in GLES2/gl2.h:94
	DST_ALPHA = 0x0304
	// ONE_MINUS_DST_ALPHA as defined in GLES2/gl2.h:95
	ONE_MINUS_DST_ALPHA = 0x0305
	// DST_COLOR as defined in GLES2/gl2.h:96
	DST_COLOR = 0x0306
	// ONE_MINUS_DST_COLOR as defined in GLES2/gl2.h:97
	ONE_MINUS_DST_COLOR = 0x0307
	// SRC_ALPHA_SATURATE as defined in GLES2/gl2.h:98
	SRC_ALPHA_SATURATE = 0x0308
	// FUNC_ADD as defined in GLES2/gl2.h:99
	FUNC_ADD = 0x8006
	// BLEND_EQUATION as defined in GLES2/gl2.h:100
	BLEND_EQUATION = 0x8009
	// BLEND_EQUATION_RGB as defined in GLES2/gl2.h:101
	BLEND_EQUATION_RGB = 0x8009
	// BLEND_EQUATION_ALPHA as defined in GLES2/gl2.h:102
	BLEND_EQUATION_ALPHA = 0x883D
	// FUNC_SUBTRACT as defined in GLES2/gl2.h:103
	FUNC_SUBTRACT = 0x800A
	// FUNC_REVERSE_SUBTRACT as defined in GLES2/gl2.h:104
	FUNC_REVERSE_SUBTRACT = 0x800B
	// BLEND_DST_RGB as defined in GLES2/gl2.h:105
	BLEND_DST_RGB = 0x80C8
	// BLEND_SRC_RGB as defined in GLES2/gl2.h:106
	BLEND_SRC_RGB = 0x80C9
	// BLEND_DST_ALPHA as defined in GLES2/gl2.h:107
	BLEND_DST_ALPHA = 0x80CA
	// BLEND_SRC_ALPHA as defined in GLES2/gl2.h:108
	BLEND_SRC_ALPHA = 0x80CB
	// CONSTANT_COLOR as defined in GLES2/gl2.h:109
	CONSTANT_COLOR = 0x8001
	// ONE_MINUS_CONSTANT_COLOR as defined in GLES2/gl2.h:110
	ONE_MINUS_CONSTANT_COLOR = 0x8002
	// CONSTANT_ALPHA as defined in GLES2/gl2.h:111
	CONSTANT_ALPHA = 0x8003
	// ONE_MINUS_CONSTANT_ALPHA as defined in GLES2/gl2.h:112
	ONE_MINUS_CONSTANT_ALPHA = 0x8004
	// BLEND_COLOR as defined in GLES2/gl2.h:113
	BLEND_COLOR = 0x8005
	// ARRAY_BUFFER as defined in GLES2/gl2.h:114
	ARRAY_BUFFER = 0x8892
	// ELEMENT_ARRAY_BUFFER as defined in GLES2/gl2.h:115
	ELEMENT_ARRAY_BUFFER = 0x8893
	// ARRAY_BUFFER_BINDING as defined in GLES2/gl2.h:116
	ARRAY_BUFFER_BINDING = 0x8894
	// ELEMENT_ARRAY_BUFFER_BINDING as defined in GLES2/gl2.h:117
	ELEMENT_ARRAY_BUFFER_BINDING = 0x8895
	// STREAM_DRAW as defined in GLES2/gl2.h:118
	STREAM_DRAW = 0x88E0
	// STATIC_DRAW as defined in GLES2/gl2.h:119
	STATIC_DRAW = 0x88E4
	// DYNAMIC_DRAW as defined in GLES2/gl2.h:120
	DYNAMIC_DRAW = 0x88E8
	// BUFFER_SIZE as defined in GLES2/gl2.h:121
	BUFFER_SIZE = 0x8764
	// BUFFER_USAGE as defined in GLES2/gl2.h:122
	BUFFER_USAGE = 0x8765
	// CURRENT_VERTEX_ATTRIB as defined in GLES2/gl2.h:123
	CURRENT_VERTEX_ATTRIB = 0x8626
	// FRONT as defined in GLES2/gl2.h:124
	FRONT = 0x0404
	// BACK as defined in GLES2/gl2.h:125
	BACK = 0x0405
	// FRONT_AND_BACK as defined in GLES2/gl2.h:126
	FRONT_AND_BACK = 0x0408
	// TEXTURE_2D as defined in GLES2/gl2.h:127
	TEXTURE_2D = 0x0DE1
	// CULL_FACE as defined in GLES2/gl2.h:128
	CULL_FACE = 0x0B44
	// BLEND as defined in GLES2/gl2.h:129
	BLEND = 0x0BE2
	// DITHER as defined in GLES2/gl2.h:130
	DITHER = 0x0BD0
	// STENCIL_TEST as defined in GLES2/gl2.h:131
	STENCIL_TEST = 0x0B90
	// DEPTH_TEST as defined in GLES2/gl2.h:132
	DEPTH_TEST = 0x0B71
	// SCISSOR_TEST as defined in GLES2/gl2.h:133
	SCISSOR_TEST = 0x0C11
	// POLYGON_OFFSET_FILL as defined in GLES2/gl2.h:134
	POLYGON_OFFSET_FILL = 0x8037
	// SAMPLE_ALPHA_TO_COVERAGE as defined in GLES2/gl2.h:135
	SAMPLE_ALPHA_TO_COVERAGE = 0x809E
	// SAMPLE_COVERAGE as defined in GLES2/gl2.h:136
	SAMPLE_COVERAGE = 0x80A0
	// NO_ERROR as defined in GLES2/gl2.h:137
	NO_ERROR = 0
	// INVALID_ENUM as defined in GLES2/gl2.h:138
	INVALID_ENUM = 0x0500
	// INVALID_VALUE as defined in GLES2/gl2.h:139
	INVALID_VALUE = 0x0501
	// INVALID_OPERATION as defined in GLES2/gl2.h:140
	INVALID_OPERATION = 0x0502
	// OUT_OF_MEMORY as defined in GLES2/gl2.h:141
	OUT_OF_MEMORY = 0x0505
	// CW as defined in GLES2/gl2.h:142
	CW = 0x0900
	// CCW as defined in GLES2/gl2.h:143
	CCW = 0x0901
	// LINE_WIDTH as defined in GLES2/gl2.h:144
	LINE_WIDTH = 0x0B21
	// ALIASED_POINT_SIZE_RANGE as defined in GLES2/gl2.h:145
	ALIASED_POINT_SIZE_RANGE = 0x846D
	// ALIASED_LINE_WIDTH_RANGE as defined in GLES2/gl2.h:146
	ALIASED_LINE_WIDTH_RANGE = 0x846E
	// CULL_FACE_MODE as defined in GLES2/gl2.h:147
	CULL_FACE_MODE = 0x0B45
	// FRONT_FACE as defined in GLES2/gl2.h:148
	FRONT_FACE = 0x0B46
	// DEPTH_RANGE as defined in GLES2/gl2.h:149
	DEPTH_RANGE = 0x0B70
	// DEPTH_WRITEMASK as defined in GLES2/gl2.h:150
	DEPTH_WRITEMASK = 0x0B72
	// DEPTH_CLEAR_VALUE as defined in GLES2/gl2.h:151
	DEPTH_CLEAR_VALUE = 0x0B73
	// DEPTH_FUNC as defined in GLES2/gl2.h:152
	DEPTH_FUNC = 0x0B74
	// STENCIL_CLEAR_VALUE as defined in GLES2/gl2.h:153
	STENCIL_CLEAR_VALUE = 0x0B91
	// STENCIL_FUNC as defined in GLES2/gl2.h:154
	STENCIL_FUNC = 0x0B92
	// STENCIL_FAIL as defined in GLES2/gl2.h:155
	STENCIL_FAIL = 0x0B94
	// STENCIL_PASS_DEPTH_FAIL as defined in GLES2/gl2.h:156
	STENCIL_PASS_DEPTH_FAIL = 0x0B95
	// STENCIL_PASS_DEPTH_PASS as defined in GLES2/gl2.h:157
	STENCIL_PASS_DEPTH_PASS = 0x0B96
	// STENCIL_REF as defined in GLES2/gl2.h:158
	STENCIL_REF = 0x0B97
	// STENCIL_VALUE_MASK as defined in GLES2/gl2.h:159
	STENCIL_VALUE_MASK = 0x0B93
	// STENCIL_WRITEMASK as defined in GLES2/gl2.h:160
	STENCIL_WRITEMASK = 0x0B98
	// STENCIL_BACK_FUNC as defined in GLES2/gl2.h:161
	STENCIL_BACK_FUNC = 0x8800
	// STENCIL_BACK_FAIL as defined in GLES2/gl2.h:162
	STENCIL_BACK_FAIL = 0x8801
	// STENCIL_BACK_PASS_DEPTH_FAIL as defined in GLES2/gl2.h:163
	STENCIL_BACK_PASS_DEPTH_FAIL = 0x8802
	// STENCIL_BACK_PASS_DEPTH_PASS as defined in GLES2/gl2.h:164
	STENCIL_BACK_PASS_DEPTH_PASS = 0x8803
	// STENCIL_BACK_REF as defined in GLES2/gl2.h:165
	STENCIL_BACK_REF = 0x8CA3
	// STENCIL_BACK_VALUE_MASK as defined in GLES2/gl2.h:166
	STENCIL_BACK_VALUE_MASK = 0x8CA4
	// STENCIL_BACK_WRITEMASK as defined in GLES2/gl2.h:167
	STENCIL_BACK_WRITEMASK = 0x8CA5
	// VIEWPORT as defined in GLES2/gl2.h:168
	VIEWPORT = 0x0BA2
	// SCISSOR_BOX as defined in GLES2/gl2.h:169
	SCISSOR_BOX = 0x0C10
	// COLOR_CLEAR_VALUE as defined in GLES2/gl2.h:170
	COLOR_CLEAR_VALUE = 0x0C22
	// COLOR_WRITEMASK as defined in GLES2/gl2.h:171
	COLOR_WRITEMASK = 0x0C23
	// UNPACK_ALIGNMENT as defined in GLES2/gl2.h:172
	UNPACK_ALIGNMENT = 0x0CF5
	// PACK_ALIGNMENT as defined in GLES2/gl2.h:173
	PACK_ALIGNMENT = 0x0D05
	// MAX_TEXTURE_SIZE as defined in GLES2/gl2.h:174
	MAX_TEXTURE_SIZE = 0x0D33
	// MAX_VIEWPORT_DIMS as defined in GLES2/gl2.h:175
	MAX_VIEWPORT_DIMS = 0x0D3A
	// SUBPIXEL_BITS as defined in GLES2/gl2.h:176
	SUBPIXEL_BITS = 0x0D50
	// RED_BITS as defined in GLES2/gl2.h:177
	RED_BITS = 0x0D52
	// GREEN_BITS as defined in GLES2/gl2.h:178
	GREEN_BITS = 0x0D53
	// BLUE_BITS as defined in GLES2/gl2.h:179
	BLUE_BITS = 0x0D54
	// ALPHA_BITS as defined in GLES2/gl2.h:180
	ALPHA_BITS = 0x0D55
	// DEPTH_BITS as defined in GLES2/gl2.h:181
	DEPTH_BITS = 0x0D56
	// STENCIL_BITS as defined in GLES2/gl2.h:182
	STENCIL_BITS = 0x0D57
	// POLYGON_OFFSET_UNITS as defined in GLES2/gl2.h:183
	POLYGON_OFFSET_UNITS = 0x2A00
	// POLYGON_OFFSET_FACTOR as defined in GLES2/gl2.h:184
	POLYGON_OFFSET_FACTOR = 0x8038
	// TEXTURE_BINDING_2D as defined in GLES2/gl2.h:185
	TEXTURE_BINDING_2D = 0x8069
	// SAMPLE_BUFFERS as defined in GLES2/gl2.h:186
	SAMPLE_BUFFERS = 0x80A8
	// SAMPLES as defined in GLES2/gl2.h:187
	SAMPLES = 0x80A9
	// SAMPLE_COVERAGE_VALUE as defined in GLES2/gl2.h:188
	SAMPLE_COVERAGE_VALUE = 0x80AA
	// SAMPLE_COVERAGE_INVERT as defined in GLES2/gl2.h:189
	SAMPLE_COVERAGE_INVERT = 0x80AB
	// NUM_COMPRESSED_TEXTURE_FORMATS as defined in GLES2/gl2.h:190
	NUM_COMPRESSED_TEXTURE_FORMATS = 0x86A2
	// COMPRESSED_TEXTURE_FORMATS as defined in GLES2/gl2.h:191
	COMPRESSED_TEXTURE_FORMATS = 0x86A3
	// DONT_CARE as defined in GLES2/gl2.h:192
	DONT_CARE = 0x1100
	// FASTEST as defined in GLES2/gl2.h:193
	FASTEST = 0x1101
	// NICEST as defined in GLES2/gl2.h:194
	NICEST = 0x1102
	// GENERATE_MIPMAP_HINT as defined in GLES2/gl2.h:195
	GENERATE_MIPMAP_HINT = 0x8192
	// BYTE as defined in GLES2/gl2.h:196
	BYTE = 0x1400
	// UNSIGNED_BYTE as defined in GLES2/gl2.h:197
	UNSIGNED_BYTE = 0x1401
	// SHORT as defined in GLES2/gl2.h:198
	SHORT = 0x1402
	// UNSIGNED_SHORT as defined in GLES2/gl2.h:199
	UNSIGNED_SHORT = 0x1403
	// INT as defined in GLES2/gl2.h:200
	INT = 0x1404
	// UNSIGNED_INT as defined in GLES2/gl2.h:201
	UNSIGNED_INT = 0x1405
	// FLOAT as defined in GLES2/gl2.h:202
	FLOAT = 0x1406
	// FIXED as defined in GLES2/gl2.h:203
	FIXED = 0x140C
	// DEPTH_COMPONENT as defined in GLES2/gl2.h:204
	DEPTH_COMPONENT = 0x1902
	// ALPHA as defined in GLES2/gl2.h:205
	ALPHA = 0x1906
	// RGB as defined in GLES2/gl2.h:206
	RGB = 0x1907
	// RGBA as defined in GLES2/gl2.h:207
	RGBA = 0x1908
	// LUMINANCE as defined in GLES2/gl2.h:208
	LUMINANCE = 0x1909
	// LUMINANCE_ALPHA as defined in GLES2/gl2.h:209
	LUMINANCE_ALPHA = 0x190A
	// UNSIGNED_SHORT_4_4_4_4 as defined in GLES2/gl2.h:210
	UNSIGNED_SHORT_4_4_4_4 = 0x8033
	// UNSIGNED_SHORT_5_5_5_1 as defined in GLES2/gl2.h:211
	UNSIGNED_SHORT_5_5_5_1 = 0x8034
	// UNSIGNED_SHORT_5_6_5 as defined in GLES2/gl2.h:212
	UNSIGNED_SHORT_5_6_5 = 0x8363
	// FRAGMENT_SHADER as defined in GLES2/gl2.h:213
	FRAGMENT_SHADER = 0x8B30
	// VERTEX_SHADER as defined in GLES2/gl2.h:214
	VERTEX_SHADER = 0x8B31
	// MAX_VERTEX_ATTRIBS as defined in GLES2/gl2.h:215
	MAX_VERTEX_ATTRIBS = 0x8869
	// MAX_VERTEX_UNIFORM_VECTORS as defined in GLES2/gl2.h:216
	MAX_VERTEX_UNIFORM_VECTORS = 0x8DFB
	// MAX_VARYING_VECTORS as defined in GLES2/gl2.h:217
	MAX_VARYING_VECTORS = 0x8DFC
	// MAX_COMBINED_TEXTURE_IMAGE_UNITS as defined in GLES2/gl2.h:218
	MAX_COMBINED_TEXTURE_IMAGE_UNITS = 0x8B4D
	// MAX_VERTEX_TEXTURE_IMAGE_UNITS as defined in GLES2/gl2.h:219
	MAX_VERTEX_TEXTURE_IMAGE_UNITS = 0x8B4C
	// MAX_TEXTURE_IMAGE_UNITS as defined in GLES2/gl2.h:220
	MAX_TEXTURE_IMAGE_UNITS = 0x8872
	// MAX_FRAGMENT_UNIFORM_VECTORS as defined in GLES2/gl2.h:221
	MAX_FRAGMENT_UNIFORM_VECTORS = 0x8DFD
	// SHADER_TYPE as defined in GLES2/gl2.h:222
	SHADER_TYPE = 0x8B4F
	// DELETE_STATUS as defined in GLES2/gl2.h:223
	DELETE_STATUS = 0x8B80
	// LINK_STATUS as defined in GLES2/gl2.h:224
	LINK_STATUS = 0x8B82
	// VALIDATE_STATUS as defined in GLES2/gl2.h:225
	VALIDATE_STATUS = 0x8B83
	// ATTACHED_SHADERS as defined in GLES2/gl2.h:226
	ATTACHED_SHADERS = 0x8B85
	// ACTIVE_UNIFORMS as defined in GLES2/gl2.h:227
	ACTIVE_UNIFORMS = 0x8B86
	// ACTIVE_UNIFORM_MAX_LENGTH as defined in GLES2/gl2.h:228
	ACTIVE_UNIFORM_MAX_LENGTH = 0x8B87
	// ACTIVE_ATTRIBUTES as defined in GLES2/gl2.h:229
	ACTIVE_ATTRIBUTES = 0x8B89
	// ACTIVE_ATTRIBUTE_MAX_LENGTH as defined in GLES2/gl2.h:230
	ACTIVE_ATTRIBUTE_MAX_LENGTH = 0x8B8A
	// SHADING_LANGUAGE_VERSION as defined in GLES2/gl2.h:231
	SHADING_LANGUAGE_VERSION = 0x8B8C
	// CURRENT_PROGRAM as defined in GLES2/gl2.h:232
	CURRENT_PROGRAM = 0x8B8D
	// NEVER as defined in GLES2/gl2.h:233
	NEVER = 0x0200
	// LESS as defined in GLES2/gl2.h:234
	LESS = 0x0201
	// EQUAL as defined in GLES2/gl2.h:235
	EQUAL = 0x0202
	// LEQUAL as defined in GLES2/gl2.h:236
	LEQUAL = 0x0203
	// GREATER as defined in GLES2/gl2.h:237
	GREATER = 0x0204
	// NOTEQUAL as defined in GLES2/gl2.h:238
	NOTEQUAL = 0x0205
	// GEQUAL as defined in GLES2/gl2.h:239
	GEQUAL = 0x0206
	// ALWAYS as defined in GLES2/gl2.h:240
	ALWAYS = 0x0207
	// KEEP as defined in GLES2/gl2.h:241
	KEEP = 0x1E00
	// REPLACE as defined in GLES2/gl2.h:242
	REPLACE = 0x1E01
	// INCR as defined in GLES2/gl2.h:243
	INCR = 0x1E02
	// DECR as defined in GLES2/gl2.h:244
	DECR = 0x1E03
	// INVERT as defined in GLES2/gl2.h:245
	INVERT = 0x150A
	// INCR_WRAP as defined in GLES2/gl2.h:246
	INCR_WRAP = 0x8507
	// DECR_WRAP as defined in GLES2/gl2.h:247
	DECR_WRAP = 0x8508
	// VENDOR as defined in GLES2/gl2.h:248
	VENDOR = 0x1F00
	// RENDERER as defined in GLES2/gl2.h:249
	RENDERER = 0x1F01
	// VERSION as defined in GLES2/gl2.h:250
	VERSION = 0x1F02
	// EXTENSIONS as defined in GLES2/gl2.h:251
	EXTENSIONS = 0x1F03
	// NEAREST as defined in GLES2/gl2.h:252
	NEAREST = 0x2600
	// LINEAR as defined in GLES2/gl2.h:253
	LINEAR = 0x2601
	// NEAREST_MIPMAP_NEAREST as defined in GLES2/gl2.h:254
	NEAREST_MIPMAP_NEAREST = 0x2700
	// LINEAR_MIPMAP_NEAREST as defined in GLES2/gl2.h:255
	LINEAR_MIPMAP_NEAREST = 0x2701
	// NEAREST_MIPMAP_LINEAR as defined in GLES2/gl2.h:256
	NEAREST_MIPMAP_LINEAR = 0x2702
	// LINEAR_MIPMAP_LINEAR as defined in GLES2/gl2.h:257
	LINEAR_MIPMAP_LINEAR = 0x2703
	// TEXTURE_MAG_FILTER as defined in GLES2/gl2.h:258
	TEXTURE_MAG_FILTER = 0x2800
	// TEXTURE_MIN_FILTER as defined in GLES2/gl2.h:259
	TEXTURE_MIN_FILTER = 0x2801
	// TEXTURE_WRAP_S as defined in GLES2/gl2.h:260
	TEXTURE_WRAP_S = 0x2802
	// TEXTURE_WRAP_T as defined in GLES2/gl2.h:261
	TEXTURE_WRAP_T = 0x2803
	// TEXTURE as defined in GLES2/gl2.h:262
	TEXTURE = 0x1702
	// TEXTURE_CUBE_MAP as defined in GLES2/gl2.h:263
	TEXTURE_CUBE_MAP = 0x8513
	// TEXTURE_BINDING_CUBE_MAP as defined in GLES2/gl2.h:264
	TEXTURE_BINDING_CUBE_MAP = 0x8514
	// TEXTURE_CUBE_MAP_POSITIVE_X as defined in GLES2/gl2.h:265
	TEXTURE_CUBE_MAP_POSITIVE_X = 0x8515
	// TEXTURE_CUBE_MAP_NEGATIVE_X as defined in GLES2/gl2.h:266
	TEXTURE_CUBE_MAP_NEGATIVE_X = 0x8516
	// TEXTURE_CUBE_MAP_POSITIVE_Y as defined in GLES2/gl2.h:267
	TEXTURE_CUBE_MAP_POSITIVE_Y = 0x8517
	// TEXTURE_CUBE_MAP_NEGATIVE_Y as defined in GLES2/gl2.h:268
	TEXTURE_CUBE_MAP_NEGATIVE_Y = 0x8518
	// TEXTURE_CUBE_MAP_POSITIVE_Z as defined in GLES2/gl2.h:269
	TEXTURE_CUBE_MAP_POSITIVE_Z = 0x8519
	// TEXTURE_CUBE_MAP_NEGATIVE_Z as defined in GLES2/gl2.h:270
	TEXTURE_CUBE_MAP_NEGATIVE_Z = 0x851A
	// MAX_CUBE_MAP_TEXTURE_SIZE as defined in GLES2/gl2.h:271
	MAX_CUBE_MAP_TEXTURE_SIZE = 0x851C
	// TEXTURE0 as defined in GLES2/gl2.h:272
	TEXTURE0 = 0x84C0
	// TEXTURE1 as defined in GLES2/gl2.h:273
	TEXTURE1 = 0x84C1
	// TEXTURE2 as defined in GLES2/gl2.h:274
	TEXTURE2 = 0x84C2
	// TEXTURE3 as defined in GLES2/gl2.h:275
	TEXTURE3 = 0x84C3
	// TEXTURE4 as defined in GLES2/gl2.h:276
	TEXTURE4 = 0x84C4
	// TEXTURE5 as defined in GLES2/gl2.h:277
	TEXTURE5 = 0x84C5
	// TEXTURE6 as defined in GLES2/gl2.h:278
	TEXTURE6 = 0x84C6
	// TEXTURE7 as defined in GLES2/gl2.h:279
	TEXTURE7 = 0x84C7
	// TEXTURE8 as defined in GLES2/gl2.h:280
	TEXTURE8 = 0x84C8
	// TEXTURE9 as defined in GLES2/gl2.h:281
	TEXTURE9 = 0x84C9
	// TEXTURE10 as defined in GLES2/gl2.h:282
	TEXTURE10 = 0x84CA
	// TEXTURE11 as defined in GLES2/gl2.h:283
	TEXTURE11 = 0x84CB
	// TEXTURE12 as defined in GLES2/gl2.h:284
	TEXTURE12 = 0x84CC
	// TEXTURE13 as defined in GLES2/gl2.h:285
	TEXTURE13 = 0x84CD
	// TEXTURE14 as defined in GLES2/gl2.h:286
	TEXTURE14 = 0x84CE
	// TEXTURE15 as defined in GLES2/gl2.h:287
	TEXTURE15 = 0x84CF
	// TEXTURE16 as defined in GLES2/gl2.h:288
	TEXTURE16 = 0x84D0
	// TEXTURE17 as defined in GLES2/gl2.h:289
	TEXTURE17 = 0x84D1
	// TEXTURE18 as defined in GLES2/gl2.h:290
	TEXTURE18 = 0x84D2
	// TEXTURE19 as defined in GLES2/gl2.h:291
	TEXTURE19 = 0x84D3
	// TEXTURE20 as defined in GLES2/gl2.h:292
	TEXTURE20 = 0x84D4
	// TEXTURE21 as defined in GLES2/gl2.h:293
	TEXTURE21 = 0x84D5
	// TEXTURE22 as defined in GLES2/gl2.h:294
	TEXTURE22 = 0x84D6
	// TEXTURE23 as defined in GLES2/gl2.h:295
	TEXTURE23 = 0x84D7
	// TEXTURE24 as defined in GLES2/gl2.h:296
	TEXTURE24 = 0x84D8
	// TEXTURE25 as defined in GLES2/gl2.h:297
	TEXTURE25 = 0x84D9
	// TEXTURE26 as defined in GLES2/gl2.h:298
	TEXTURE26 = 0x84DA
	// TEXTURE27 as defined in GLES2/gl2.h:299
	TEXTURE27 = 0x84DB
	// TEXTURE28 as defined in GLES2/gl2.h:300
	TEXTURE28 = 0x84DC
	// TEXTURE29 as defined in GLES2/gl2.h:301
	TEXTURE29 = 0x84DD
	// TEXTURE30 as defined in GLES2/gl2.h:302
	TEXTURE30 = 0x84DE
	// TEXTURE31 as defined in GLES2/gl2.h:303
	TEXTURE31 = 0x84DF
	// ACTIVE_TEXTURE as defined in GLES2/gl2.h:304
	ACTIVE_TEXTURE = 0x84E0
	// REPEAT as defined in GLES2/gl2.h:305
	REPEAT = 0x2901
	// CLAMP_TO_EDGE as defined in GLES2/gl2.h:306
	CLAMP_TO_EDGE = 0x812F
	// MIRRORED_REPEAT as defined in GLES2/gl2.h:307
	MIRRORED_REPEAT = 0x8370
	// FLOAT_VEC2 as defined in GLES2/gl2.h:308
	FLOAT_VEC2 = 0x8B50
	// FLOAT_VEC3 as defined in GLES2/gl2.h:309
	FLOAT_VEC3 = 0x8B51
	// FLOAT_VEC4 as defined in GLES2/gl2.h:310
	FLOAT_VEC4 = 0x8B52
	// INT_VEC2 as defined in GLES2/gl2.h:311
	INT_VEC2 = 0x8B53
	// INT_VEC3 as defined in GLES2/gl2.h:312
	INT_VEC3 = 0x8B54
	// INT_VEC4 as defined in GLES2/gl2.h:313
	INT_VEC4 = 0x8B55
	// BOOL as defined in GLES2/gl2.h:314
	BOOL = 0x8B56
	// BOOL_VEC2 as defined in GLES2/gl2.h:315
	BOOL_VEC2 = 0x8B57
	// BOOL_VEC3 as defined in GLES2/gl2.h:316
	BOOL_VEC3 = 0x8B58
	// BOOL_VEC4 as defined in GLES2/gl2.h:317
	BOOL_VEC4 = 0x8B59
	// FLOAT_MAT2 as defined in GLES2/gl2.h:318
	FLOAT_MAT2 = 0x8B5A
	// FLOAT_MAT3 as defined in GLES2/gl2.h:319
	FLOAT_MAT3 = 0x8B5B
	// FLOAT_MAT4 as defined in GLES2/gl2.h:320
	FLOAT_MAT4 = 0x8B5C
	// SAMPLER_2D as defined in GLES2/gl2.h:321
	SAMPLER_2D = 0x8B5E
	// SAMPLER_CUBE as defined in GLES2/gl2.h:322
	SAMPLER_CUBE = 0x8B60
	// VERTEX_ATTRIB_ARRAY_ENABLED as defined in GLES2/gl2.h:323
	VERTEX_ATTRIB_ARRAY_ENABLED = 0x8622
	// VERTEX_ATTRIB_ARRAY_SIZE as defined in GLES2/gl2.h:324
	VERTEX_ATTRIB_ARRAY_SIZE = 0x8623
	// VERTEX_ATTRIB_ARRAY_STRIDE as defined in GLES2/gl2.h:325
	VERTEX_ATTRIB_ARRAY_STRIDE = 0x8624
	// VERTEX_ATTRIB_ARRAY_TYPE as defined in GLES2/gl2.h:326
	VERTEX_ATTRIB_ARRAY_TYPE = 0x8625
	// VERTEX_ATTRIB_ARRAY_NORMALIZED as defined in GLES2/gl2.h:327
	VERTEX_ATTRIB_ARRAY_NORMALIZED = 0x886A
	// VERTEX_ATTRIB_ARRAY_POINTER as defined in GLES2/gl2.h:328
	VERTEX_ATTRIB_ARRAY_POINTER = 0x8645
	// VERTEX_ATTRIB_ARRAY_BUFFER_BINDING as defined in GLES2/gl2.h:329
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = 0x889F
	// IMPLEMENTATION_COLOR_READ_TYPE as defined in GLES2/gl2.h:330
	IMPLEMENTATION_COLOR_READ_TYPE = 0x8B9A
	// IMPLEMENTATION_COLOR_READ_FORMAT as defined in GLES2/gl2.h:331
	IMPLEMENTATION_COLOR_READ_FORMAT = 0x8B9B
	// COMPILE_STATUS as defined in GLES2/gl2.h:332
	COMPILE_STATUS = 0x8B81
	// INFO_LOG_LENGTH as defined in GLES2/gl2.h:333
	INFO_LOG_LENGTH = 0x8B84
	// SHADER_SOURCE_LENGTH as defined in GLES2/gl2.h:334
	SHADER_SOURCE_LENGTH = 0x8B88
	// SHADER_COMPILER as defined in GLES2/gl2.h:335
	SHADER_COMPILER = 0x8DFA
	// SHADER_BINARY_FORMATS as defined in GLES2/gl2.h:336
	SHADER_BINARY_FORMATS = 0x8DF8
	// NUM_SHADER_BINARY_FORMATS as defined in GLES2/gl2.h:337
	NUM_SHADER_BINARY_FORMATS = 0x8DF9
	// LOW_FLOAT as defined in GLES2/gl2.h:338
	LOW_FLOAT = 0x8DF0
	// MEDIUM_FLOAT as defined in GLES2/gl2.h:339
	MEDIUM_FLOAT = 0x8DF1
	// HIGH_FLOAT as defined in GLES2/gl2.h:340
	HIGH_FLOAT = 0x8DF2
	// LOW_INT as defined in GLES2/gl2.h:341
	LOW_INT = 0x8DF3
	// MEDIUM_INT as defined in GLES2/gl2.h:342
	MEDIUM_INT = 0x8DF4
	// HIGH_INT as defined in GLES2/gl2.h:343
	HIGH_INT = 0x8DF5
	// FRAMEBUFFER as defined in GLES2/gl2.h:344
	FRAMEBUFFER = 0x8D40
	// RENDERBUFFER as defined in GLES2/gl2.h:345
	RENDERBUFFER = 0x8D41
	// RGBA4 as defined in GLES2/gl2.h:346
	RGBA4 = 0x8056
	// RGB5_A1 as defined in GLES2/gl2.h:347
	RGB5_A1 = 0x8057
	// RGB565 as defined in GLES2/gl2.h:348
	RGB565 = 0x8D62
	// DEPTH_COMPONENT16 as defined in GLES2/gl2.h:349
	DEPTH_COMPONENT16 = 0x81A5
	// STENCIL_INDEX8 as defined in GLES2/gl2.h:350
	STENCIL_INDEX8 = 0x8D48
	// RENDERBUFFER_WIDTH as defined in GLES2/gl2.h:351
	RENDERBUFFER_WIDTH = 0x8D42
	// RENDERBUFFER_HEIGHT as defined in GLES2/gl2.h:352
	RENDERBUFFER_HEIGHT = 0x8D43
	// RENDERBUFFER_INTERNAL_FORMAT as defined in GLES2/gl2.h:353
	RENDERBUFFER_INTERNAL_FORMAT = 0x8D44
	// RENDERBUFFER_RED_SIZE as defined in GLES2/gl2.h:354
	RENDERBUFFER_RED_SIZE = 0x8D50
	// RENDERBUFFER_GREEN_SIZE as defined in GLES2/gl2.h:355
	RENDERBUFFER_GREEN_SIZE = 0x8D51
	// RENDERBUFFER_BLUE_SIZE as defined in GLES2/gl2.h:356
	RENDERBUFFER_BLUE_SIZE = 0x8D52
	// RENDERBUFFER_ALPHA_SIZE as defined in GLES2/gl2.h:357
	RENDERBUFFER_ALPHA_SIZE = 0x8D53
	// RENDERBUFFER_DEPTH_SIZE as defined in GLES2/gl2.h:358
	RENDERBUFFER_DEPTH_SIZE = 0x8D54
	// RENDERBUFFER_STENCIL_SIZE as defined in GLES2/gl2.h:359
	RENDERBUFFER_STENCIL_SIZE = 0x8D55
	// FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE as defined in GLES2/gl2.h:360
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE = 0x8CD0
	// FRAMEBUFFER_ATTACHMENT_OBJECT_NAME as defined in GLES2/gl2.h:361
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME = 0x8CD1
	// FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL as defined in GLES2/gl2.h:362
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL = 0x8CD2
	// FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE as defined in GLES2/gl2.h:363
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = 0x8CD3
	// COLOR_ATTACHMENT0 as defined in GLES2/gl2.h:364
	COLOR_ATTACHMENT0 = 0x8CE0
	// DEPTH_ATTACHMENT as defined in GLES2/gl2.h:365
	DEPTH_ATTACHMENT = 0x8D00
	// STENCIL_ATTACHMENT as defined in GLES2/gl2.h:366
	STENCIL_ATTACHMENT = 0x8D20
	// NONE as defined in GLES2/gl2.h:367
	NONE = 0
	// FRAMEBUFFER_COMPLETE as defined in GLES2/gl2.h:368
	FRAMEBUFFER_COMPLETE = 0x8CD5
	// FRAMEBUFFER_INCOMPLETE_ATTACHMENT as defined in GLES2/gl2.h:369
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT = 0x8CD6
	// FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT as defined in GLES2/gl2.h:370
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = 0x8CD7
	// FRAMEBUFFER_INCOMPLETE_DIMENSIONS as defined in GLES2/gl2.h:371
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS = 0x8CD9
	// FRAMEBUFFER_UNSUPPORTED as defined in GLES2/gl2.h:372
	FRAMEBUFFER_UNSUPPORTED = 0x8CDD
	// FRAMEBUFFER_BINDING as defined in GLES2/gl2.h:373
	FRAMEBUFFER_BINDING = 0x8CA6
	// RENDERBUFFER_BINDING as defined in GLES2/gl2.h:374
	RENDERBUFFER_BINDING = 0x8CA7
	// MAX_RENDERBUFFER_SIZE as defined in GLES2/gl2.h:375
	MAX_RENDERBUFFER_SIZE = 0x84E8
	// INVALID_FRAMEBUFFER_OPERATION as defined in GLES2/gl2.h:376
	INVALID_FRAMEBUFFER_OPERATION = 0x0506
)
