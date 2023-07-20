export namespace backend {
	
	export class Address {
	    street: string;
	    postcode: string;
	
	    static createFrom(source: any = {}) {
	        return new Address(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.street = source["street"];
	        this.postcode = source["postcode"];
	    }
	}
	export class Book {
	    name: string;
	    author: string;
	    desc: string;
	    sites: models.BookSite[];
	
	    static createFrom(source: any = {}) {
	        return new Book(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.author = source["author"];
	        this.desc = source["desc"];
	        this.sites = this.convertValues(source["sites"], models.BookSite);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Person {
	    name: string;
	    age: number;
	    address?: Address;
	
	    static createFrom(source: any = {}) {
	        return new Person(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.age = source["age"];
	        this.address = this.convertValues(source["address"], Address);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace models {
	
	export class Book {
	    author: string;
	    bookName: string;
	    bookId: string;
	
	    static createFrom(source: any = {}) {
	        return new Book(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.author = source["author"];
	        this.bookName = source["bookName"];
	        this.bookId = source["bookId"];
	    }
	}
	export class BookSiteContentUri {
	    uri: string;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new BookSiteContentUri(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uri = source["uri"];
	        this.content = source["content"];
	    }
	}
	export class BookSiteChapterUri {
	    uri: string;
	    chapterList: string;
	    chapterName: string;
	    chapterUrl: string;
	    chapterId: string;
	
	    static createFrom(source: any = {}) {
	        return new BookSiteChapterUri(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uri = source["uri"];
	        this.chapterList = source["chapterList"];
	        this.chapterName = source["chapterName"];
	        this.chapterUrl = source["chapterUrl"];
	        this.chapterId = source["chapterId"];
	    }
	}
	export class BookSiteUri {
	    uri: string;
	    bookList: string;
	    author: string;
	    bookName: string;
	    bookUri: string;
	    bookId: string;
	
	    static createFrom(source: any = {}) {
	        return new BookSiteUri(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uri = source["uri"];
	        this.bookList = source["bookList"];
	        this.author = source["author"];
	        this.bookName = source["bookName"];
	        this.bookUri = source["bookUri"];
	        this.bookId = source["bookId"];
	    }
	}
	export class BookSite {
	    name: string;
	    uri: string;
	    searchUrl: BookSiteUri;
	    chapterUrl: BookSiteChapterUri;
	    contentUrl: BookSiteContentUri;
	    checkStep: number;
	
	    static createFrom(source: any = {}) {
	        return new BookSite(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.uri = source["uri"];
	        this.searchUrl = this.convertValues(source["searchUrl"], BookSiteUri);
	        this.chapterUrl = this.convertValues(source["chapterUrl"], BookSiteChapterUri);
	        this.contentUrl = this.convertValues(source["contentUrl"], BookSiteContentUri);
	        this.checkStep = source["checkStep"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	export class Chapter {
	    bookId: string;
	    chapterId: string;
	    chapterName: string;
	
	    static createFrom(source: any = {}) {
	        return new Chapter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bookId = source["bookId"];
	        this.chapterId = source["chapterId"];
	        this.chapterName = source["chapterName"];
	    }
	}
	export class Content {
	    bookId: string;
	    chapterId: string;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new Content(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bookId = source["bookId"];
	        this.chapterId = source["chapterId"];
	        this.content = source["content"];
	    }
	}

}

