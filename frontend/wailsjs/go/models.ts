export namespace db {
	
	export class CatboxAuth {
	
	
	    static createFrom(source: any = {}) {
	        return new CatboxAuth(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class Queries {
	
	
	    static createFrom(source: any = {}) {
	        return new Queries(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class UpdateCatboxAuthParams {
	
	
	    static createFrom(source: any = {}) {
	        return new UpdateCatboxAuthParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

export namespace sql {
	
	export class Tx {
	
	
	    static createFrom(source: any = {}) {
	        return new Tx(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

