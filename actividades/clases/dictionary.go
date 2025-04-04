package clases 

// struct Dictionary usando map
type Dictionary struct {
    data map[string]string
}

// creacion nuevo dictionary 
func NewDictionary() *Dictionary {
    return &Dictionary{
        data: make(map[string]string),
    }
}

// añade un nuevo par key-value
func (d *Dictionary) Add(key, value string) {
    d.data[key] = value
}

// regresa el valor asociado a la llave key
// regresa un booleano si la llave existe o no
func (d *Dictionary) Get(key string) (string, bool) {
    value, exists := d.data[key]
    return value, exists
}

// borra un par ket-value 
func (d *Dictionary) Remove(key string) {
    delete(d.data, key)
}

// actualiza el valor asociado a la llave key
// regresa un booleano si la llave existe o no
func (d *Dictionary) Update(key, value string) bool {
    if _, exists := d.data[key]; exists {
        d.data[key] = value
        return true
    }
    return false
}

// revisa si la llave key existe en el dictionario
func (d *Dictionary) Contains(key string) bool {
    _, exists := d.data[key]
    return exists
}

// regresa el numero de elementos en el dictionario
func (d *Dictionary) Size() int {
    return len(d.data)
}

//	borra todos los elementos en el dictionario
func (d *Dictionary) Clear() {
    d.data = make(map[string]string)
}

// regresa todo el dictionario
func (d *Dictionary) GetAll() map[string]string {
    return d.data
}
