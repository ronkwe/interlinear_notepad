export namespace core {
	
	export class DocConfig {
	    line_count: number;
	    line_labels: string[];
	
	    static createFrom(source: any = {}) {
	        return new DocConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.line_count = source["line_count"];
	        this.line_labels = source["line_labels"];
	    }
	}
	export class GlossColumn {
	    lines: string[];
	
	    static createFrom(source: any = {}) {
	        return new GlossColumn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.lines = source["lines"];
	    }
	}
	export class GlossBlock {
	    columns: GlossColumn[];
	
	    static createFrom(source: any = {}) {
	        return new GlossBlock(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.columns = this.convertValues(source["columns"], GlossColumn);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	
	export class GlossDocument {
	    version: string;
	    config: DocConfig;
	    blocks: GlossBlock[];
	
	    static createFrom(source: any = {}) {
	        return new GlossDocument(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.config = this.convertValues(source["config"], DocConfig);
	        this.blocks = this.convertValues(source["blocks"], GlossBlock);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

