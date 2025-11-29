export namespace core {
	
	export class LineOption {
	    label: string;
	    visible: boolean;
	
	    static createFrom(source: any = {}) {
	        return new LineOption(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.label = source["label"];
	        this.visible = source["visible"];
	    }
	}
	export class DocConfig {
	    line_count: number;
	    line_options: LineOption[];
	    show_labels: boolean;
	
	    static createFrom(source: any = {}) {
	        return new DocConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.line_count = source["line_count"];
	        this.line_options = this.convertValues(source["line_options"], LineOption);
	        this.show_labels = source["show_labels"];
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
	export class GlossColumn {
	    id: string;
	    lines: string[];
	
	    static createFrom(source: any = {}) {
	        return new GlossColumn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
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

