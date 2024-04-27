export namespace main {
	
	export class Xor {
	    first: string;
	    second: string;
	    third: string;
	    four: string;
	
	    static createFrom(source: any = {}) {
	        return new Xor(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.first = source["first"];
	        this.second = source["second"];
	        this.third = source["third"];
	        this.four = source["four"];
	    }
	}

}

