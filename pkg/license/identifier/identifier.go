// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package identifier

// Canonical names of the licenses.
// The names come from the https://spdx.org/licenses website, and are
// also the filenames of the licenses in licenseclassifier/licenses.
const (
	AFL11                       = "AFL-1.1"
	AFL12                       = "AFL-1.2"
	AFL20                       = "AFL-2.0"
	AFL21                       = "AFL-2.1"
	AFL30                       = "AFL-3.0"
	AGPL10                      = "AGPL-1.0"
	AGPL30                      = "AGPL-3.0"
	Apache10                    = "Apache-1.0"
	Apache11                    = "Apache-1.1"
	Apache20                    = "Apache-2.0"
	APSL10                      = "APSL-1.0"
	APSL11                      = "APSL-1.1"
	APSL12                      = "APSL-1.2"
	APSL20                      = "APSL-2.0"
	Artistic10cl8               = "Artistic-1.0-cl8"
	Artistic10Perl              = "Artistic-1.0-Perl"
	Artistic10                  = "Artistic-1.0"
	Artistic20                  = "Artistic-2.0"
	BCL                         = "BCL"
	Beerware                    = "Beerware"
	BSD2ClauseFreeBSD           = "BSD-2-Clause-FreeBSD"
	BSD2ClauseNetBSD            = "BSD-2-Clause-NetBSD"
	BSD2Clause                  = "BSD-2-Clause"
	BSD3ClauseAttribution       = "BSD-3-Clause-Attribution"
	BSD3ClauseClear             = "BSD-3-Clause-Clear"
	BSD3ClauseLBNL              = "BSD-3-Clause-LBNL"
	BSD3Clause                  = "BSD-3-Clause"
	BSD4Clause                  = "BSD-4-Clause"
	BSD4ClauseUC                = "BSD-4-Clause-UC"
	BSDProtection               = "BSD-Protection"
	BSL10                       = "BSL-1.0"
	CC010                       = "CC0-1.0"
	CCBY10                      = "CC-BY-1.0"
	CCBY20                      = "CC-BY-2.0"
	CCBY25                      = "CC-BY-2.5"
	CCBY30                      = "CC-BY-3.0"
	CCBY40                      = "CC-BY-4.0"
	CCBYNC10                    = "CC-BY-NC-1.0"
	CCBYNC20                    = "CC-BY-NC-2.0"
	CCBYNC25                    = "CC-BY-NC-2.5"
	CCBYNC30                    = "CC-BY-NC-3.0"
	CCBYNC40                    = "CC-BY-NC-4.0"
	CCBYNCND10                  = "CC-BY-NC-ND-1.0"
	CCBYNCND20                  = "CC-BY-NC-ND-2.0"
	CCBYNCND25                  = "CC-BY-NC-ND-2.5"
	CCBYNCND30                  = "CC-BY-NC-ND-3.0"
	CCBYNCND40                  = "CC-BY-NC-ND-4.0"
	CCBYNCSA10                  = "CC-BY-NC-SA-1.0"
	CCBYNCSA20                  = "CC-BY-NC-SA-2.0"
	CCBYNCSA25                  = "CC-BY-NC-SA-2.5"
	CCBYNCSA30                  = "CC-BY-NC-SA-3.0"
	CCBYNCSA40                  = "CC-BY-NC-SA-4.0"
	CCBYND10                    = "CC-BY-ND-1.0"
	CCBYND20                    = "CC-BY-ND-2.0"
	CCBYND25                    = "CC-BY-ND-2.5"
	CCBYND30                    = "CC-BY-ND-3.0"
	CCBYND40                    = "CC-BY-ND-4.0"
	CCBYSA10                    = "CC-BY-SA-1.0"
	CCBYSA20                    = "CC-BY-SA-2.0"
	CCBYSA25                    = "CC-BY-SA-2.5"
	CCBYSA30                    = "CC-BY-SA-3.0"
	CCBYSA40                    = "CC-BY-SA-4.0"
	CDDL10                      = "CDDL-1.0"
	CDDL11                      = "CDDL-1.1"
	CommonsClause               = "Commons-Clause"
	CPAL10                      = "CPAL-1.0"
	CPL10                       = "CPL-1.0"
	EGenix                      = "eGenix"
	EPL10                       = "EPL-1.0"
	EUPL10                      = "EUPL-1.0"
	EUPL11                      = "EUPL-1.1"
	Facebook2Clause             = "Facebook-2-Clause"
	Facebook3Clause             = "Facebook-3-Clause"
	FacebookExamples            = "Facebook-Examples"
	FreeImage                   = "FreeImage"
	FTL                         = "FTL"
	GPL10                       = "GPL-1.0"
	GPL20                       = "GPL-2.0"
	GPL20withautoconfexception  = "GPL-2.0-with-autoconf-exception"
	GPL20withbisonexception     = "GPL-2.0-with-bison-exception"
	GPL20withclasspathexception = "GPL-2.0-with-classpath-exception"
	GPL20withfontexception      = "GPL-2.0-with-font-exception"
	GPL20withGCCexception       = "GPL-2.0-with-GCC-exception"
	GPL30                       = "GPL-3.0"
	GPL30withautoconfexception  = "GPL-3.0-with-autoconf-exception"
	GPL30withGCCexception       = "GPL-3.0-with-GCC-exception"
	GUSTFont                    = "GUST-Font-License"
	ImageMagick                 = "ImageMagick"
	IPL10                       = "IPL-1.0"
	ISC                         = "ISC"
	LGPL20                      = "LGPL-2.0"
	LGPL21                      = "LGPL-2.1"
	LGPL30                      = "LGPL-3.0"
	LGPLLR                      = "LGPLLR"
	Libpng                      = "Libpng"
	Lil10                       = "Lil-1.0"
	LPL102                      = "LPL-1.02"
	LPL10                       = "LPL-1.0"
	LPPL13c                     = "LPPL-1.3c"
	MIT                         = "MIT"
	MPL10                       = "MPL-1.0"
	MPL11                       = "MPL-1.1"
	MPL20                       = "MPL-2.0"
	MSPL                        = "MS-PL"
	NCSA                        = "NCSA"
	NPL10                       = "NPL-1.0"
	NPL11                       = "NPL-1.1"
	OFL                         = "OFL"
	OpenSSL                     = "OpenSSL"
	OSL10                       = "OSL-1.0"
	OSL11                       = "OSL-1.1"
	OSL20                       = "OSL-2.0"
	OSL21                       = "OSL-2.1"
	OSL30                       = "OSL-3.0"
	PHP301                      = "PHP-3.01"
	PHP30                       = "PHP-3.0"
	PIL                         = "PIL"
	Python20complete            = "Python-2.0-complete"
	Python20                    = "Python-2.0"
	QPL10                       = "QPL-1.0"
	Ruby                        = "Ruby"
	SGIB10                      = "SGI-B-1.0"
	SGIB11                      = "SGI-B-1.1"
	SGIB20                      = "SGI-B-2.0"
	SISSL12                     = "SISSL-1.2"
	SISSL                       = "SISSL"
	Sleepycat                   = "Sleepycat"
	UnicodeTOU                  = "Unicode-TOU"
	Unlicense                   = "Unlicense"
	W3C19980720                 = "W3C-19980720"
	W3C                         = "W3C"
	WTFPL                       = "WTFPL"
	X11                         = "X11"
	Xnet                        = "Xnet"
	Zend20                      = "Zend-2.0"
	ZlibAcknowledgement         = "zlib-acknowledgement"
	Zlib                        = "Zlib"
	ZPL11                       = "ZPL-1.1"
	ZPL20                       = "ZPL-2.0"
	ZPL21                       = "ZPL-2.1"
)

// List returns license identifier in slice
func List() []string {
	return []string{
		AFL11,
		AFL12,
		AFL20,
		AFL21,
		AFL30,
		AGPL10,
		AGPL30,
		Apache10,
		Apache11,
		Apache20,
		APSL10,
		APSL11,
		APSL12,
		APSL20,
		Artistic10cl8,
		Artistic10Perl,
		Artistic10,
		Artistic20,
		BCL,
		Beerware,
		BSD2ClauseFreeBSD,
		BSD2ClauseNetBSD,
		BSD2Clause,
		BSD3ClauseAttribution,
		BSD3ClauseClear,
		BSD3ClauseLBNL,
		BSD3Clause,
		BSD4Clause,
		BSD4ClauseUC,
		BSDProtection,
		BSL10,
		CC010,
		CCBY10,
		CCBY20,
		CCBY25,
		CCBY30,
		CCBY40,
		CCBYNC10,
		CCBYNC20,
		CCBYNC25,
		CCBYNC30,
		CCBYNC40,
		CCBYNCND10,
		CCBYNCND20,
		CCBYNCND25,
		CCBYNCND30,
		CCBYNCND40,
		CCBYNCSA10,
		CCBYNCSA20,
		CCBYNCSA25,
		CCBYNCSA30,
		CCBYNCSA40,
		CCBYND10,
		CCBYND20,
		CCBYND25,
		CCBYND30,
		CCBYND40,
		CCBYSA10,
		CCBYSA20,
		CCBYSA25,
		CCBYSA30,
		CCBYSA40,
		CDDL10,
		CDDL11,
		CommonsClause,
		CPAL10,
		CPL10,
		EGenix,
		EPL10,
		EUPL10,
		EUPL11,
		Facebook2Clause,
		Facebook3Clause,
		FacebookExamples,
		FreeImage,
		FTL,
		GPL10,
		GPL20,
		GPL20withautoconfexception,
		GPL20withbisonexception,
		GPL20withclasspathexception,
		GPL20withfontexception,
		GPL20withGCCexception,
		GPL30,
		GPL30withautoconfexception,
		GPL30withGCCexception,
		GUSTFont,
		ImageMagick,
		IPL10,
		ISC,
		LGPL20,
		LGPL21,
		LGPL30,
		LGPLLR,
		Libpng,
		Lil10,
		LPL102,
		LPL10,
		LPPL13c,
		MIT,
		MPL10,
		MPL11,
		MPL20,
		MSPL,
		NCSA,
		NPL10,
		NPL11,
		OFL,
		OpenSSL,
		OSL10,
		OSL11,
		OSL20,
		OSL21,
		OSL30,
		PHP301,
		PHP30,
		PIL,
		Python20complete,
		Python20,
		QPL10,
		Ruby,
		SGIB10,
		SGIB11,
		SGIB20,
		SISSL12,
		SISSL,
		Sleepycat,
		UnicodeTOU,
		Unlicense,
		W3C19980720,
		W3C,
		WTFPL,
		X11,
		Xnet,
		Zend20,
		ZlibAcknowledgement,
		Zlib,
		ZPL11,
		ZPL20,
		ZPL21,
	}
}
